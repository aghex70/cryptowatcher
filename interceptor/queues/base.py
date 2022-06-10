from abc import ABC, abstractmethod
from typing import Dict, List


class QueueCommon(ABC):
    @abstractmethod
    def get_queue_by_name(self, queue_name: str) -> Dict:
        raise NotImplementedError

    @abstractmethod
    def get_queue_by_url(self, queue_url: str) -> Dict:
        raise NotImplementedError

    @abstractmethod
    def get_queue_url(self, queue_name: str) -> str:
        raise NotImplementedError

    @abstractmethod
    def get_queues(self, queue_prefix: str) -> List:
        raise NotImplementedError

    @abstractmethod
    def delete_message(self, queue_url: str, **kwargs) -> None:
        raise NotImplementedError

    @abstractmethod
    def delete_messages(self, queue_url: str, **kwargs) -> Dict:
        raise NotImplementedError
