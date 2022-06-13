from typing import Dict, Union

from botocore.exceptions import ClientError
from config.logger import logger
from provider.aws.sqs.base import SQSCommon
from queues.receiver import QueueReceiver


class SQSReceiver(QueueReceiver, SQSCommon):
    def __init__(self, client, resource=None):
        super().__init__(client, resource)

    def receive_message(self, queue_url: str) -> Dict:
        """Receive SQS message
        :param queue_url: SQS queue url
        :return: response
        """
        logger.info("Retrieving message. Url: %s", queue_url)
        try:
            response = self.client.receive_message(
                QueueUrl=queue_url, MaxNumberOfMessages=1
            )
            logger.info("Received message. Url: %s", queue_url)
        except ClientError as exc:
            logger.error("Error receiving message. Url: %s", queue_url)
            raise exc
        return response

    def receive_messages(self, queue_url: str, **kwargs) -> Union[Dict, None]:
        """Receive SQS messages
        :param queue_url: SQS queue url
        :kwargs: attributes
        :return: response
        """
        logger.info("Retrieving messages. Url: %s", queue_url)
        messages_number = kwargs["messages_number"]
        if messages_number > 10:
            logger.fatal("Maximum number of messages is 10")
        try:
            response = self.client.receive_message(
                QueueUrl=queue_url, MaxNumberOfMessages=messages_number
            )
            trades = response.get("Messages")
            if not trades:
                logger.warning("No messages received from queue. ")
                return

            messages = response["Messages"]
            logger.info("Received messages. Url: %s", queue_url)
        except ClientError as exc:
            logger.error("Error receiving messages. Url: %s", queue_url)
            raise exc
        return messages
