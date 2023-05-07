import json

import redis.asyncio as aioredis
from aiogram import Bot
from src.consumer.resolve import resolve_message


async def consume(redis: aioredis.Redis, bot: Bot):
    pubsub = redis.pubsub()

    await pubsub.subscribe("telegram")

    while True:
        message = await pubsub.get_message(ignore_subscribe_messages=True)
        if message is not None:
            mess, user_ids = parse_message(message["data"])
            await resolve_message(redis, bot, mess, user_ids)


def parse_message(message: str):
    data = json.loads(message)

    return data.get("Message"), data.get("UserIds")
