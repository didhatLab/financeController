import abc
import datetime
from typing import Protocol, Any

from src.models.currency import CurrencyRate


class CurrentCurrencyRatesRepository(Protocol):
    @abc.abstractmethod
    async def save_new_rate(self, rate: CurrencyRate) -> Any:
        pass


class CurrencyHistoryRepository(Protocol):
    @abc.abstractmethod
    async def get_rates_for_time(self, time: datetime.datetime) -> CurrencyRate:
        pass

    @abc.abstractmethod
    async def save_to_history_rate(self, rate: CurrencyRate) -> Any:
        pass
