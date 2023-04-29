import asyncio
import datetime as dt
import time

import aiohttp
import redis.asyncio as aioredis
from motor.motor_asyncio import AsyncIOMotorClient, AsyncIOMotorDatabase
from scheduler import Scheduler

from src.app.application import UpdaterCurrencyApplication


async def main():
    client = AsyncIOMotorClient("localhost:27017")
    db: AsyncIOMotorDatabase = client.get_database("currency")

    redis = aioredis.Redis()
    session = aiohttp.ClientSession()

    app = UpdaterCurrencyApplication(db, redis, session)

    service_for_run = app.build()
    schedule = Scheduler()
    await service_for_run.update_currencies()
    schedule.cyclic(dt.timedelta(hours=1), service_for_run.update_currencies)

    while True:
        schedule.exec_jobs()
        time.sleep(10)


if __name__ == "__main__":
    asyncio.run(main())
