from abc import ABC, abstractmethod
from typing import Dict


class Consumer(ABC):
    @abstractmethod
    def subscribe(self, queue_url: str) -> Dict:
        raise NotImplementedError

    @abstractmethod
    def unsubscribe(self):
        raise NotImplementedError
