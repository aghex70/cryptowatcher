import json
from typing import Dict, List, Union

from botocore.exceptions import ClientError
from config.logger import logger
from provider.aws.sqs.base import SQSCommon
from queues.sender import QueueSender


class SQSSender(QueueSender, SQSCommon):
    def __init__(self, client, resource=None):
        super().__init__(client, resource)

    def create_queue(self, queue_name: str, **kwargs) -> str:
        """Create SQS queue
        :queue_name: SQS queue name
        :kwargs: attributes
        :return: queue url
        """
        logger.info("Creating queue. Name: %s, kwargs: %s", queue_name, kwargs)
        try:
            response = self.client.create_queue(
                QueueName=queue_name,
                Attributes=kwargs.get("attributes", {}),
                tags=kwargs.get("tags", {}),
            )
            queue_url = response["QueueUrl"]
            logger.info("Created queue. Name: %s, url: %s", queue_name, queue_url)
        except ClientError as exc:
            logger.error(
                "Error creating queue. Name: %s, configuration: %s", queue_name, kwargs
            )
            raise exc
        return queue_url

    def configure_queue(self, queue_url: str, **kwargs) -> None:
        """Configure SQS queue
        :param queue_url: SQS queue url
        :param kwargs: attributes
        """
        logger.info("Configuring queue. Url: %s", queue_url)
        try:
            self.client.set_queue_attributes(QueueUrl=queue_url, Attributes=kwargs)
            logger.info(
                "Configured queue. Url: %s, attributes: %s",
                queue_url,
                kwargs,
            )
        except ClientError as exc:
            logger.error(
                "Error configuring queue. Url: %s, attributes: %s", queue_url, kwargs
            )
            raise exc

    def purge_queue(self, queue_url: str) -> None:
        """Deletes all the messages in a queue
        :param queue_url: SQS queue url
        """
        logger.info("Purging queue. Url: %s", queue_url)
        try:
            self.client.purge_queue(QueueUrl=queue_url)
            logger.info("Purged queue. Url: %s", queue_url)
        except ClientError as exc:
            logger.error("Error purging queue. Url: %s", queue_url)
            raise exc

    def remove_queue(self, queue_url: str) -> None:
        """Deletes the queue
        :param queue_url: SQS queue url
        """
        logger.info("Deleting queue. Url: %s", queue_url)
        try:
            self.client.delete_queue(QueueUrl=queue_url)
            logger.info("Deleted queue. Url: %s", queue_url)
        except ClientError as exc:
            logger.error("Error deleting queue. Url: %s", queue_url)
            raise exc

    def send_message(self, queue_url: str, message: Union[Dict, str]) -> Dict:
        """Send a message to SQS queue
        :param queue_url: SQS queue url
        :param message: message
        :return: response
        """
        if isinstance(message, dict):
            message = json.dumps(message)

        logger.info(
            "Sending message. Url: %s, message: %s",
            queue_url,
            message,
        )
        try:
            response = self.client.send_message(QueueUrl=queue_url, MessageBody=message)
            logger.info(
                "Sent message. Url: %s, message: %s",
                queue_url,
                message,
            )
        except ClientError as exc:
            logger.error(
                "Error sending message. Url: %s, message: %s",
                queue_url,
                message,
            )
            raise exc
        return response

    def send_messages(self, queue_url: str, messages: List) -> Dict:
        """Send multiple messages to SQS queue
        :param queue_url: SQS queue url
        :param messages: messages
        :return: response
        """
        logger.info("Sending messages. Url: %s", queue_url)
        if not isinstance(messages, list):
            logger.fatal(
                "Error sending messages. Messages must be a list. Messages: %s",
                messages,
            )
        try:
            response = self.client.send_message_batch(
                QueueUrl=queue_url, Entries=messages
            )
            logger.info("Sent messages. Url: %s", queue_url)
        except ClientError as exc:
            logger.error("Error sending messages. Url: %s", queue_url)
            raise exc
        logger.info(response)
        return response
