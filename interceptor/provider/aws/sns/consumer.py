import logging
from typing import Dict

from botocore.exceptions import ClientError
from messages.consumer import Consumer
from provider.aws.sqs.base import SQSCommon

logger = logging.getLogger(__name__)


class SNSConsumer(Consumer):
    def __init__(self, client, resource=None):
        self.client = client
        self.resource = resource

    def subscribe(self, queue_url: str) -> Dict:
        logger.info("Creating topic")
        try:
            response = self.client.create_topic()
        except ClientError as exc:
            logger.error("Error creating topic")
            raise exc

        logger.info("Created topic")
        return response

    def unsubscribe(self):
        logger.info("Creating topic")
        try:
            response = self.client.create_topic()
        except ClientError as exc:
            logger.error("Error creating topic")
            raise exc

        logger.info("Created topic")
        return response
