"""Main module for the cryptowatcher courier."""
import time

import click
from app.interceptor import Interceptor
from config import settings
from config.logger import logger
from provider.aws.client import AWSClient
from provider.aws.resource import AWSResource
from provider.aws.sqs.receiver import SQSReceiver

RECEIVER = "receiver"


@click.command()
@click.option("--queue", type=click.Choice([RECEIVER], case_sensitive=False))
def main(queue: str) -> None:
    logger.info("Executing main")
    if queue == RECEIVER:
        receiver = run_receiver()
        interceptor = Interceptor(receiver=receiver)
        # while True:
        logger.info("Starting to read trades from queue")
        response = interceptor.receive_trades()


def run_receiver():
    client = AWSClient(service_name=settings.AWS_SQS_SERVICE_NAME).create()
    resource = AWSResource(service_name=settings.AWS_SQS_SERVICE_NAME).create()
    sender = SQSReceiver(client=client, resource=resource)
    return sender


if __name__ == "__main__":
    main()
