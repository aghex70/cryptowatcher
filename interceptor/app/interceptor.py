"""Module in charge of reading the orders queue and stored them in the database."""
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
        messages_kwargs = {"messages_number": 2}
        received = False
        while not received:
            trades = self.receiver.receive_messages(self.queue_url, **messages_kwargs)
            if not trades:
                logger.info("No trades received. Retrying in 5 seconds...")
                continue

            received = True
            logger.info("Received trades. Trades: %s", trades)
            self.persist_trades(trades)
        saved_trades = Trade.get_trades()
        logger.info("Saved trades. Trades: %s", saved_trades)

    def persist_trades(self, trades: List):
        for trade in trades:
            self.persist_trade(trade)
            trade = Trade.create(**trade)
            trade.save()

    def persist_trade(self, trade: Dict):
        receipt_handle = trade.pop("ReceiptHandle", "")
        logger.info("Persisting trade in database. Trade: %s", trade)
        message_kwargs = {"receipt_handle": receipt_handle}
        self.receiver.delete_message(self.queue_url, **message_kwargs)
