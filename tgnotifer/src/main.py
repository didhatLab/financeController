import asyncio

import redis.asyncio as aioredis
from aiogram import Bot, Dispatcher
from aiogram.fsm.storage.redis import RedisStorage

from src.routes.start import start_route
from src.routes.setting import notification_setting_route
from src.routes.login import login_menu_router
from src.consumer.main import consume
from src.config import get_bot_token, get_redis_config


async def main():
    redis_config = get_redis_config()

    redis = aioredis.Redis(
        host=redis_config.host,
        port=redis_config.port,
        db=redis_config.db,
        password=redis_config.password,
    )
    storage = RedisStorage(redis)
    disp = Dispatcher(storage=storage)

    disp.include_routers(
        start_route(), notification_setting_route(redis), login_menu_router(redis)
    )

    bot = Bot(get_bot_token())
    asyncio.create_task(consume(redis, bot))
    await disp.start_polling(bot)


if __name__ == "__main__":
    asyncio.run(main())
