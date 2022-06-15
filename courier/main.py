"""Main module for the cryptowatcher courier."""
import time

import click
from app.courier import Courier
from config import settings
from config.logger import logger
from provider.aws.client import AWSClient
from provider.aws.resource import AWSResource
from provider.aws.sqs.sender import SQSSender

SENDER = "sender"


@click.command()
@click.option("--queue", type=click.Choice([SENDER], case_sensitive=False))
def main(queue: str) -> None:
    logger.info("Executing main")
    count, sent = 0, False
    if queue == SENDER:
        sender = run_sender()
        courier = Courier(sender=sender)
        while not sent and count < 10:
            trades = courier.get_trades(limit=10)
            if not trades:
                logger.warning("No trades found. Retrying in 5 seconds...")
                time.sleep(5)
                continue

            logger.info("Got trades: %s", trades)
            courier.send_trades(trades=trades)
            sent = True

        if not sent:
            logger.warning("No trades sent. Shutting down...")
            return



def run_sender():
    client = AWSClient(service_name=settings.AWS_SQS_SERVICE_NAME).create()
    resource = AWSResource(service_name=settings.AWS_SQS_SERVICE_NAME).create()
    sender = SQSSender(client=client, resource=resource)
    return sender


if __name__ == "__main__":
    main()
