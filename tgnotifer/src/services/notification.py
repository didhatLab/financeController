import aiohttp
import json
import redis.asyncio as aioredis


async def _set_notify_status(token: str, status: str):
    async with aiohttp.ClientSession() as session:
        async with session.post(
            "http://localhost:4005/set",
            json={"Status": status},
            headers={"Auth-Token": token}
        ):
            pass


async def turn_off_notify(redis: aioredis.Redis, chat_id: int):
    token = await get_user_token(redis, chat_id)

    await _set_notify_status(token, "inactive")


async def turn_on_notify(redis: aioredis.Redis, chat_id: int):
    token = await get_user_token(redis, chat_id)

    await _set_notify_status(token, "active")


async def get_user_token(redis: aioredis.Redis, chat_id: int) -> str:
    info = await redis.get(f"{str(chat_id)}:chat")

    data = json.loads(info)

    return data.get("token")
