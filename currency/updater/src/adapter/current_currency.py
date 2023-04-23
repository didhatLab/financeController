import json
import redis.asyncio as aioredis
from typing import Any

from src.models.currency import CurrencyRate


class CurrentCurrencyRatesRepositoryImp:
    def __init__(self, redis: aioredis.Redis, currency_rate_key: str):
        self._redis = redis
        self._redis_key_for_rate = currency_rate_key

    async def save_new_rate(self, rate: CurrencyRate) -> Any:
        rates_json = json.loads(rate.to_dict())

        await self._redis.set(self._redis_key_for_rate, rates_json)
