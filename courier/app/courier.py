"""Module in charge of retrieving orders from the database and sending them to the queue."""
import json
from typing import Dict, List

from app.repo import get_collection, get_database
from bson.objectid import ObjectId
from config import settings
from config.logger import logger


class Courier:
    def __init__(self, sender):
        self.database = get_database()
        self.collection = get_collection(self.database)
        self.sender = sender

    def get_trades(self, limit: int = 25):
        """Get trades from database
        :param limit: limit of trades
        """
        logger.info("Retrieving trades. Limit: %s", limit)
        return list(self.collection.find({}).limit(limit))

    def delete_trade(self, trade_id: int):
        """Delete trade from database
        :param trade_id: trade id
        """
        logger.info("Deleting trade. Id: %s", trade_id)
        return self.collection.delete_one({"_id": trade_id})

    def delete_trades(self, ids_list: List):
        """Delete trades from database
        :param ids_list: list of trade ids
        """
        logger.info("Deleting trades. Ids: %s", ids_list)
        return self.collection.delete_many(filter={"_id": {"$in": ids_list}})

    def retrieve_queue(self):
        """Retrieve queue"""
        queue = self.sender.get_queue_by_name()
        if not queue:
            queue = self.sender.create_queue(settings.QUEUE_NAME)
        return queue

    def send_trades(self, trades: List):
        """Send trades to queue
        :param trades: list of trades
        """
        queue_url = self.retrieve_queue()
        wrapped_trades = self.build_messages(trades)
        logger.info(wrapped_trades)
        response = self.sender.send_messages(queue_url, wrapped_trades)
        delivered_trades = response.get("Successful", [])
        logger.info(
            "Proceeding to remove delivered trades from database: %s", delivered_trades
        )
        self.remove_delivered_trades(delivered_trades)

    def remove_delivered_trades(self, trades: List):
        """Remove trades from database
        :param trades: list of trades
        """
        delivered_ids = [ObjectId(trade["Id"]) for trade in trades]
        self.delete_trades(delivered_ids)

    def build_messages(self, trades: List) -> List:
        """Build messages from trades
        :param trades: list of trades
        :return: list of messages, list of trade ids
        """
        wrapped_messages = [self.build_message(trade) for trade in trades]
        return wrapped_messages

    @staticmethod
    def build_message(trade: Dict) -> Dict:
        """Build message from trade
        :param trade: trade
        :return: message
        """
        trade_id = trade.pop("_id")
        wrapped_message = {"Id": str(trade_id), "MessageBody": json.dumps(trade)}
        return wrapped_message
