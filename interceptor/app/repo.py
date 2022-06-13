import enum
import json
from typing import List

from config import settings
from config.logger import logger
from sqlalchemy import (
    BigInteger,
    Column,
    DateTime,
    Enum,
    Float,
    ForeignKey,
    Integer,
    Numeric,
    String,
    create_engine,
    func,
    text,
)
from sqlalchemy.exc import OperationalError
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import Session, relationship, sessionmaker


def start_database_engine():
    """
    Start the database engine
    :return: The database engine
    """
    return create_engine(settings.DATABASE_URI, encoding="utf-8")


def start_session(engine) -> Session:
    """
    Start a database session
    :param engine: The database engine
    :return: A database session
    """
    return sessionmaker(bind=engine, autocommit=True)()


def create_table_base(engine):
    """
    Create the base table
    :param engine: The database engine
    """
    return declarative_base(bind=engine)


def ping(session: Session) -> True:
    """
    Ping the database by executing a simple query
    :param session: The database session
    :return: True if the database is reachable
    """
    logger.info("Pinging database. URI: %s", settings.DATABASE_URI)
    logger.info("PING (database).")
    try:
        session.execute(text("SELECT 1"))
        logger.info("PONG (database)")
        return True
    except OperationalError as e:
        logger.error("Database connection error: %s", e)
        raise


def create_tables(engine):
    """
    Create the tables
    :param engine: The database engine
    """
    logger.info("Creating tables. URI: %s", settings.DATABASE_URI)
    Base.metadata.create_all(bind=engine)


database_engine = start_database_engine()
database_session = start_session(database_engine)
ping(database_session)
Base = create_table_base(database_engine)


class Trade(Base):
    """
    The trade table
    """

    __tablename__ = "trades"

    class Currency(enum.Enum):
        EUR = 1
        USD = 2
        PESETAS = 3

    class TradeType(enum.Enum):
        BUY = 1
        SELL = 2

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    provider_id = Column(ForeignKey("providers.id"), nullable=False)
    trade_id = Column(BigInteger, index=True)
    symbols_from = Column(String(10), nullable=False)
    symbols_to = Column(String(10), nullable=False)
    currency = Column(Enum(Currency), default=Currency.USD, nullable=False)
    price = Column(Float(precision=10), nullable=False)
    quantity = Column(Float(precision=10), nullable=False)
    trade_type = Column(Enum(TradeType), nullable=False)
    event_time = Column(DateTime, nullable=True)
    trade_time = Column(DateTime, nullable=True)
    buyer_id = Column(ForeignKey("users.id"), nullable=False)
    seller_id = Column(ForeignKey("users.id"), nullable=False)
    created_at = Column(DateTime, nullable=False, server_default=func.now())
    updated_at = Column(
        DateTime, nullable=False, server_default=func.now(), onupdate=func.now()
    )
    provider = relationship("Provider")
    buyer = relationship("User", foreign_keys=[buyer_id])
    seller = relationship("User", foreign_keys=[seller_id])

    @classmethod
    def create(cls, **kwargs) -> "Trade":
        """
        Create a trade
        :return: The created trade
        """
        logger.info("Creating trade. Trade: %s", kwargs)
        trade_info = json.loads(kwargs["Body"])

        source = trade_info["source"]
        seller_external_id = trade_info["a"]
        buyer_external_id = trade_info["b"]
        trade_id = trade_info["t"]
        symbols_to = trade_info["s"][:3]
        symbols_from = trade_info["s"][3:]
        price = float(trade_info["p"])
        quantity = float(trade_info["q"])
        trade_type = Trade.TradeType.SELL if trade_info["m"] else Trade.TradeType.BUY

        # Retrieve provider
        provider = Provider.get_or_create(name=source)
        # Retrieve buyer
        buyer = User.get_or_create(
            external_id=buyer_external_id, provider_id=provider.id
        )
        # Retrieve seller
        seller = User.get_or_create(
            external_id=seller_external_id, provider_id=provider.id
        )

        trade = Trade()
        trade.trade_id = trade_id
        trade.symbols_to = symbols_to
        trade.symbols_from = symbols_from
        trade.price = price
        trade.quantity = quantity
        trade.trade_type = trade_type
        trade.buyer_id = buyer.id
        trade.seller_id = seller.id
        trade.provider_id = provider.id
        trade.save()
        logger.info("Created trade. Trade_id: %s", trade.id)

        UserProviderRelation.get_or_create(
            user_id=buyer.id,
            user_external_id=buyer_external_id,
            provider_id=provider.id,
        )
        UserProviderRelation.get_or_create(
            user_id=seller.id,
            user_external_id=seller_external_id,
            provider_id=provider.id,
        )

        return trade

    def save(self):
        """
        Save the trade to the database
        """
        database_session.add(self)
        database_session.flush()

    @classmethod
    def get_trades(cls) -> List:
        """
        Get all trades
        :return: A list of trades
        """
        return database_session.query(cls).all()


class Provider(Base):
    """
    The provider table
    """

    __tablename__ = "providers"

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    name = Column(String(50), nullable=False)
    created_at = Column(DateTime, nullable=False, server_default=func.now())
    updated_at = Column(
        DateTime, nullable=False, server_default=func.now(), onupdate=func.now()
    )

    @classmethod
    def get_or_create(cls, name: str) -> "Provider":
        """
        Get or create the user
        """
        provider = database_session.query(cls).filter(cls.name == name).first()
        if provider:
            logger.info("Retrieved provider. Id: %s", provider.id)
            return provider

        logger.info(
            "Provider not found. Proceeding to create new provider. Name: %s", name
        )
        provider = Provider()
        provider.name = name
        provider.save()
        logger.info("Created provider. Id: %s", provider.id)
        return provider

    def save(self):
        """
        Save the provider to the database
        """
        database_session.add(self)
        database_session.flush()


class User(Base):
    """
    The user table
    """

    __tablename__ = "users"

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    external_id = Column(BigInteger, nullable=False)
    provider_id = Column(ForeignKey("providers.id"), nullable=False)
    created_at = Column(DateTime, nullable=False, server_default=func.now())
    updated_at = Column(
        DateTime, nullable=False, server_default=func.now(), onupdate=func.now()
    )

    provider = relationship("Provider")

    @classmethod
    def get_or_create(cls, external_id: int, provider_id: int) -> "User":
        """
        Get or create the user
        """
        user = (
            database_session.query(cls)
            .join(Provider)
            .filter(cls.provider_id == provider_id, cls.external_id == external_id)
            .first()
        )
        if user:
            logger.info("Retrieved user. External_id: %s", external_id)
            return user

        logger.info(
            "User not found. Proceeding to create new user. External_id: %s, provider_id: %s",
            external_id,
            provider_id,
        )

        user = User()
        user.external_id = external_id
        user.provider_id = provider_id
        user.save()
        logger.info("Created user. External_id: %s", external_id)
        return user

    def save(self):
        """
        Save the user to the database
        """
        database_session.add(self)
        database_session.flush()


class UserProviderRelation(Base):
    """
    The user provider relation table
    """

    __tablename__ = "user_provider_relations"

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    user_external_id = Column(BigInteger, nullable=False)
    user_id = Column(ForeignKey("users.id"), nullable=False)
    provider_id = Column(ForeignKey("providers.id"), nullable=False)
    created_at = Column(DateTime, nullable=False, server_default=func.now())
    updated_at = Column(
        DateTime, nullable=False, server_default=func.now(), onupdate=func.now()
    )

    @classmethod
    def get_or_create(
        cls, user_id: int, user_external_id: int, provider_id: int
    ) -> "UserProviderRelation":
        """
        Get or create the user provider relation
        :param user_id: The user id
        :param user_external_id: The user external id
        :param provider_id: The provider id
        """
        user_provider_relation = (
            database_session.query(UserProviderRelation)
            .filter(
                UserProviderRelation.user_id == user_id,
                UserProviderRelation.user_external_id == user_external_id,
                UserProviderRelation.provider_id == provider_id,
            )
            .first()
        )
        if user_provider_relation:
            logger.info(
                "Retrieved user provider relation. Id: %s",
                user_provider_relation.id,
            )
            return user_provider_relation

        logger.info(
            "User provider relation not found. Proceeding to create new relation. User_id: %s, user_external_id: %s, provider_id: %s",
            user_id,
            user_external_id,
            provider_id,
        )
        relation = UserProviderRelation()
        relation.user_external_id = user_external_id
        relation.user_id = user_id
        relation.provider_id = provider_id
        relation.save()

    def save(self):
        """
        Save the relation to the database
        """
        database_session.add(self)
        database_session.flush()


Base.metadata.create_all(bind=database_engine)
