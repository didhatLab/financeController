import asyncio
import json

from aiogram import Bot
import redis.asyncio as aioredis


async def resolve_message(
    redis: aioredis.Redis, bot: Bot, message: str, user_ids: list[int]
):
    chats_ids = []

    for uid in user_ids:
        chat_id = await get_user_chat_id(redis, uid)
        if chat_id is not None:
            chats_ids.append(chat_id)

    async with asyncio.TaskGroup() as tg:
        for ch_id in chats_ids:
            tg.create_task(bot.send_message(ch_id, message))


async def get_user_chat_id(redis: aioredis.Redis, user_id: int) -> int | None:
    chat_info = await redis.get(f"{user_id}:user")

    if chat_info is None:
        return None

    data = json.loads(chat_info)

    return int(data.get("chat_id"))
