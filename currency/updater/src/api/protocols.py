import abc
from typing import Protocol

from src.models.currency import CurrencyRate


class CurrencyAPI(Protocol):
    @abc.abstractmethod
    async def latest_currency(self) -> CurrencyRate:
        pass
