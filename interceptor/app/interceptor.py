"""Module in charge of reading the orders queue and stored them in the database."""
import json
import time
from typing import Dict, List

from app.repo import Trade
from config import settings
from config.logger import logger


class Interceptor:
    def __init__(self, receiver):
        self.receiver = receiver
        self.queue_url = self.retrieve_queue()

    def retrieve_queue(self):
        """Retrieve queue"""
        queue = self.receiver.get_queue_by_name()
        if not queue:
            queue = self.receiver.create_queue(settings.QUEUE_NAME)
        return queue

    def receive_trades(self):
        """Get trades from queue"""
        logger.info("Receiving trades")
        messages_kwargs = {"messages_number": 5}
        count, received = 0, False
        while not received and count < 10:
            trades = self.receiver.receive_messages(self.queue_url, **messages_kwargs)
            if not trades:
                logger.info("No trades received. Retrying in 5 seconds...")
                time.sleep(5)
                count += 1
                continue

            received = True
            logger.info("Received trades. Trades: %s", trades)
            self.persist_trades(trades)
        if not received:
            logger.warning("No trades received. Shutting down...")

    def persist_trades(self, trades: List):
        for trade in trades:
            self.persist_trade(trade)

    def persist_trade(self, trade: Dict):
        receipt_handle = trade.pop("ReceiptHandle", "")
        logger.info(
            "Persisting trade in database. Trade: %s, type: %s", trade, type(trade)
        )
        trade_body = trade["Body"]

        logger.info("Trade body: %s", trade_body)
        try:
            trade_body = json.loads(trade_body)
            # Only persist certain type of messages
            if "result" not in trade_body:
                Trade.create(**trade)
        except json.JSONDecodeError:
            logger.error("Error decoding trade body. Trade body: %s", trade_body)

        message_kwargs = {"receipt_handle": receipt_handle}
        self.receiver.delete_message(self.queue_url, **message_kwargs)
