import logging
from typing import Dict

from botocore.exceptions import ClientError
from messages.producer import Producer

logger = logging.getLogger(__name__)


class SNSProducer(Producer):
    def __init__(self, client, resource=None):
        self.client = client
        self.resource = resource

    def create_topic(self, queue_url: str) -> Dict:
        logger.info("Creating topic")
        try:
            response = self.client.create_topic()
        except ClientError as exc:
            logger.error("Error creating topic")
            raise exc

        logger.info("Created topic")
        return response

    def update_topic(self):
        logger.info("Creating topic")
        try:
            response = self.client.create_topic()
        except ClientError as exc:
            logger.error("Error creating topic")
            raise exc

        logger.info("Created topic")
        return response

    def delete_topic(self, queue_url: str) -> Dict:
        logger.info("Creating topic")
        try:
            response = self.client.create_topic()
        except ClientError as exc:
            logger.error("Error creating topic")
            raise exc

        logger.info("Created topic")
        return response

    def publish_message(self, queue_url: str, **kwargs) -> Dict:
        logger.info("Creating topic")
        try:
            response = self.client.create_topic()
        except ClientError as exc:
            logger.error("Error creating topic")
            raise exc

        logger.info("Created topic")
        return response

    def get_topic_configuration(self, queue_url: str, **kwargs) -> Dict:
        logger.info("Creating topic")
        try:
            response = self.client.create_topic()
        except ClientError as exc:
            logger.error("Error creating topic")
            raise exc

        logger.info("Created topic")
        return response
