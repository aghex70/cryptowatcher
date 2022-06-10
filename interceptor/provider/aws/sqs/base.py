from typing import Dict, List, Union

from botocore.exceptions import ClientError
from config import settings
from config.logger import logger
from queues.base import QueueCommon


class SQSCommon(QueueCommon):
    def __init__(
        self,
        client,
        resource=None,
    ):
        super().__init__()
        if not client:
            logger.fatal(
                "Error instantiating class. Client must be provided. Client: %s",
                client,
            )

        self.client = client
        self.resource = resource

    def get_queue_by_name(
        self, queue_name: str = settings.QUEUE_NAME
    ) -> Union[str, None]:
        """Get SQS queue by name
        :return: queue url"""
        if not self.resource:
            logger.fatal(
                "Error retrieving queues. Resource must be provided. Resource: %s",
                self.resource,
            )
        logger.info("Retrieving queue. Name: %s", queue_name)
        try:
            queue = self.resource.get_queue_by_name(
                QueueName=queue_name,
                QueueOwnerAWSAccountId=settings.AWS_ACCOUNT_ID,
            )
            logger.info("Retrieved queue. Name: %s, url: %s", queue_name, queue.url)
        except ClientError:
            logger.error("Error retrieving queue. Name: %s", queue_name)
            return

        return queue.url

    def get_queue_by_url(self, queue_url: str) -> Dict:
        """Get SQS queue by url
        :param queue_url: SQS queue url
        :return: response
        """
        logger.fatal("Error retrieving queue. Unsupported method")
        return {}

    def get_queue_url(self, queue_name: str) -> str:
        """Get SQS queue url
        :param queue_name: SQS queue name
        :return: queue url
        """
        logger.info("Retrieving queue's url. Name: %s", queue_name)
        queue_url = self.get_queue_by_name(queue_name)
        logger.info("Retrieved queue's url. Name: %s, url: %s", queue_name, queue_url)
        return queue_url

    def get_queues(self, queue_prefix: str) -> List[str]:
        """Get SQS queues
        :queue_prefix: SQS queue prefix
        :return: list of queues
        """
        if not self.resource:
            logger.fatal(
                "Error retrieving queues. Resource must be provided. Resource: %s",
                self.resource,
            )

        logger.info("Retrieving queues. Prefix: %s", queue_prefix)
        if queue_prefix:
            queue_iterator = self.resource.queues.filter(QueueNamePrefix=queue_prefix)
        else:
            queue_iterator = self.resource.queues.all()

        queues = list(queue_iterator)
        if queues:
            queues_names = ", ".join([q.url for q in queues])
            logger.info("Retrieved queues. Names: %s", queues_names)
        else:
            logger.warning(
                "No queues found with given prefix. Prefix: %s", queue_prefix
            )
        return queues

    def delete_message(self, queue_url: str, **kwargs) -> None:
        """Delete a message from SQS queue
        :param queue_url: SQS queue url
        :param kwargs: message data
        :return: response"""
        receipt_handle = kwargs["receipt_handle"]
        logger.info("Deleting message. Receipt handle: %s", receipt_handle)
        try:
            self.client.delete_message(QueueUrl=queue_url, ReceiptHandle=receipt_handle)
            logger.info(
                "Deleted message. Url: %s, receipt_handle: %s",
                queue_url,
                receipt_handle,
            )
        except ClientError as exc:
            logger.error(
                "Error deleting message. Url: %s, receipt_handle: %s",
                queue_url,
                receipt_handle,
            )
            raise exc
        return None

    def delete_messages(self, queue_url: str, **kwargs) -> Dict:
        """Delete multiple messages from SQS queue
        :param queue_url: SQS queue url
        :param kwargs: messages
        :return: response"""
        logger.info("Deleting messages. Url: %s", queue_url)
        messages = kwargs["messages"]
        if not isinstance(messages, list):
            logger.fatal(
                "Error deleting messages. Messages must be a list. Messages: %s",
                messages,
            )

        try:
            response = self.client.delete_message_batch(
                QueueUrl=queue_url, Entries=messages
            )
            logger.info("Deleted messages. Url: %s", queue_url)
        except ClientError as exc:
            logger.error("Error deleting messages. Url: %s", queue_url)
            raise exc
        return response
