from abc import abstractmethod
from typing import Dict

from queues.base import QueueCommon


class QueueReceiver(QueueCommon):
    @abstractmethod
    def receive_message(self, queue_url: str) -> Dict:
        raise NotImplementedError

    @abstractmethod
    def receive_messages(self, queue_url: str, **kwargs) -> Dict:
        raise NotImplementedError
