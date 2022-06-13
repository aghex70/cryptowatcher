from abc import ABC, abstractmethod
from typing import Dict


class Producer(ABC):
    @abstractmethod
    def create_topic(self, queue_url: str) -> Dict:
        raise NotImplementedError

    @abstractmethod
    def update_topic(self):
        raise NotImplementedError

    @abstractmethod
    def get_topic_configuration(self):
        raise NotImplementedError

    @abstractmethod
    def delete_topic(self, queue_url: str) -> Dict:
        raise NotImplementedError

    @abstractmethod
    def publish_message(self, queue_url: str, **kwargs) -> Dict:
        raise NotImplementedError

    @abstractmethod
    def publish_messages(self, queue_url: str, **kwargs) -> Dict:
        raise NotImplementedError
