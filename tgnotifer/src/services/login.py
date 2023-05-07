import json
import aiohttp
import redis.asyncio as aioredis


class LoginException(Exception):
    pass


async def login(username: str, password: str) -> tuple[str, int]:
    async with aiohttp.ClientSession() as session:
        async with session.post(
            "http://localhost:4001/token",
            json={"username": username, "password": password},
        ) as resp:
            answer = await resp.json()

    token = answer.get("Token")

    if token is None:
        raise LoginException()

    return token, int(answer.get("UserId"))


async def save_account(redis: aioredis.Redis, chat_id: int, user_id: int, token: str):
    data = {
        "user_id": user_id,
        "token": token,
    }

    await redis.set(str(chat_id), json.dumps(data))
