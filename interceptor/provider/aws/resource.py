import boto3
from config import settings


class AWSResource:
    def __init__(self, service_name: str):
        self.service_name = service_name

    def create(self):
        return boto3.resource(
            self.service_name,
            region_name=settings.AWS_REGION,
            aws_access_key_id=settings.AWS_ACCESS_KEY,
            aws_secret_access_key=settings.AWS_SECRET_KEY,
        )
