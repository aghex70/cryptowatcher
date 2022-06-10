import enum
from typing import List

from config import settings
from config.logger import logger
from sqlalchemy import (
    BigInteger,
    Column,
    DateTime,
    Enum,
    ForeignKey,
    Numeric,
    String,
    create_engine,
    func,
    text,
)
from sqlalchemy.exc import OperationalError
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import Session, sessionmaker


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
    provider = Column(ForeignKey("providers.id"), nullable=False)
    trade_id = Column(BigInteger, index=True)
    symbols_from = Column(ForeignKey("symbols.id"), nullable=False)
    symbols_to = Column(ForeignKey("symbols.id"), nullable=False)
    currency = Column(Enum(Currency), nullable=False)
    price = Column(Numeric, nullable=False)
    quantity = Column(Numeric, nullable=False)
    trade_type = Column(Enum(TradeType), nullable=False)
    event_time = Column(DateTime, nullable=False)
    trade_time = Column(DateTime, nullable=False)
    buyer = Column(ForeignKey("users.id"), nullable=False)
    seller = Column(ForeignKey("users.id"), nullable=False)
    created_at = Column(DateTime, nullable=False, server_default=func.now())
    updated_at = Column(
        DateTime, nullable=False, server_default=func.now(), onupdate=func.now()
    )

    @classmethod
    def create(cls, **kwargs) -> "Trade":
        """
        Create a trade
        :return: The created trade
        """
        logger.info("Creating trade. Trade: %s", kwargs)
        trade = Trade()
        trade.save()
        logger.info("Created trade. Trade_id: %s", trade.id)
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


class Symbol(Base):
    """
    The symbol table
    """

    __tablename__ = "symbols"

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    created_at = Column(DateTime, nullable=False, server_default=func.now())
    updated_at = Column(
        DateTime, nullable=False, server_default=func.now(), onupdate=func.now()
    )
    name = Column(String(50), nullable=False)
    symbol = Column(String(10), nullable=False)


class Provider(Base):
    """
    The provider table
    """

    __tablename__ = "providers"

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    created_at = Column(DateTime, nullable=False, server_default=func.now())
    updated_at = Column(
        DateTime, nullable=False, server_default=func.now(), onupdate=func.now()
    )
    name = Column(String(50), nullable=False)


class User(Base):
    """
    The user table
    """

    __tablename__ = "users"

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    created_at = Column(DateTime, nullable=False, server_default=func.now())
    updated_at = Column(
        DateTime, nullable=False, server_default=func.now(), onupdate=func.now()
    )
    name = Column(String(50), nullable=False)


class UserProviderRelation(Base):
    """
    The user provider relation table
    """

    __tablename__ = "user_provider_relations"

    id = Column(BigInteger, primary_key=True, autoincrement=True)
    external_id = Column(BigInteger, nullable=False)
    created_at = Column(DateTime, nullable=False, server_default=func.now())
    updated_at = Column(
        DateTime, nullable=False, server_default=func.now(), onupdate=func.now()
    )
    user = Column(ForeignKey("users.id"), nullable=False)
    provider = Column(ForeignKey("providers.id"), nullable=False)


Base.metadata.create_all(bind=database_engine)
