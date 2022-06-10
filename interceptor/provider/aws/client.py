from typing import Union

import boto3
from config import settings


class AWSClient:
    def __init__(self, service_name: str):
        self.service_name = service_name

    def create(self, token: Union[str, None] = None):
        """Login to the specified AWS service.
        :param token: The AWS session token.
        :return: The AWS service client.
        """
        if token:
            client = boto3.client(
                self.service_name,
                region_name=settings.AWS_REGION,
                aws_access_key_id=settings.AWS_ACCESS_KEY,
                aws_secret_access_key=settings.AWS_SECRET_KEY,
                aws_session_token=settings.AWS_SESSION_TOKEN,
            )
        else:
            client = boto3.client(
                self.service_name,
                region_name=settings.AWS_REGION,
                aws_access_key_id=settings.AWS_ACCESS_KEY,
                aws_secret_access_key=settings.AWS_SECRET_KEY,
            )
        return client
