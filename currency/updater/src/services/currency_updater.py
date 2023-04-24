import asyncio

from src.api.protocols import CurrencyAPI
from src.adapter.protocols import (
    CurrencyHistoryRepository,
    CurrentCurrencyRatesRepository,
)


class CurrencyUpdater:
    def __init__(
        self,
        api: CurrencyAPI,
        current_currency_repo: CurrentCurrencyRatesRepository,
        history_currency_repository: CurrencyHistoryRepository,
    ):
        self._api = api
        self._current_currency_repo = current_currency_repo
        self._history_currency_repo = history_currency_repository

    async def update_currencies(self):
        new_rate = await self._api.latest_currency()

        async with asyncio.TaskGroup() as g:
            g.create_task(self._current_currency_repo.save_new_rate(new_rate))
            g.create_task(self._history_currency_repo.save_to_history_rate(new_rate))
