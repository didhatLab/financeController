import json

import redis.asyncio as aioredis
from fastapi import FastAPI, Depends

from src.models.rate import CurrencyRate, RateForStats


app = FastAPI()


async def get_redis() -> aioredis.Redis:
    redis = aioredis.Redis()
    yield redis
    await redis.close()


@app.get("/current_currency_rate")
async def current_currency_rate(
    redis: aioredis.Redis = Depends(get_redis),
) -> CurrencyRate:
    currency_rate_json_str = await redis.get(key_for_current_rate_from_redis)
    currency_rate = json.loads(currency_rate_json_str)

    return CurrencyRate(**currency_rate)


@app.get("/currency_for_stats")
async def currency_for_stats(
    redis: aioredis.Redis = Depends(get_redis),
) -> RateForStats:
    currency_rate_json_str = await redis.get(key_for_current_rate_from_redis)
    currency_rate = CurrencyRate(**json.loads(currency_rate_json_str))

    eur = currency_rate.rates.get("RUB") * 100 / (currency_rate.rates.get("EUR") * 100)

    return RateForStats(Usd=currency_rate.rates.get("RUB"), Eur=eur)


key_for_current_rate_from_redis = "current_currency_rate"


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, port=4002)
