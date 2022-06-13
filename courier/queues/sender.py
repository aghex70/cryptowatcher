from abc import abstractmethod
from typing import Dict, Union

from queues.base import QueueCommon


class QueueSender(QueueCommon):
    @abstractmethod
    def create_queue(self, queue_name: str, **kwargs) -> str:
        raise NotImplementedError

    @abstractmethod
    def configure_queue(self, queue_url: str, **kwargs) -> None:
        raise NotImplementedError

    @abstractmethod
    def purge_queue(self, queue_url: str) -> None:
        raise NotImplementedError

    @abstractmethod
    def remove_queue(self, queue_url: str) -> None:
        raise NotImplementedError

    @abstractmethod
    def send_message(self, queue_url: str, message: Union[Dict, str]) -> Dict:
        raise NotImplementedError

    @abstractmethod
    def send_messages(self, queue_url: str, **kwargs) -> Dict:
        raise NotImplementedError
