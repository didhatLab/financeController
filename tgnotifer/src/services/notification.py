import aiohttp
import json
import redis.asyncio as aioredis


async def _set_notify_status(token: str, status: str):
    async with aiohttp.ClientSession() as session:
        async with session.post(
            "http://localhost:4005/set",
            json={"Status": status},
        ) as resp:
            answer = await resp.json()


async def turn_off_notify(redis: aioredis.Redis, chat_id: int):
    token = await get_user_token(redis, chat_id)

    await _set_notify_status(token, "inactive")


async def turn_on_notify(redis: aioredis.Redis, chat_id: int):
    token = await get_user_token(redis, chat_id)

    await _set_notify_status(token, "active")


async def get_user_token(redis: aioredis.Redis, chat_id: int) -> str:
    info = await redis.get(str(chat_id))

    data = json.loads(info)

    return data.get("token")
