import asyncio

import redis.asyncio as aioredis
from aiogram import Bot, Dispatcher
from aiogram.fsm.storage.redis import RedisStorage

from src.routes.start import start_route
from src.routes.setting import notification_setting_route
from src.routes.login import login_menu_router
from src.consumer.main import consume


async def main():
    redis = aioredis.Redis()
    storage = RedisStorage(redis)
    disp = Dispatcher(storage=storage)

    disp.include_routers(
        start_route(), notification_setting_route(redis), login_menu_router(redis)
    )

    bot = Bot("token")
    asyncio.create_task(consume(redis, bot))
    await disp.start_polling(bot)


if __name__ == "__main__":
    asyncio.run(main())
