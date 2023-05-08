import os
from typing import Final
from dataclasses import dataclass

from dotenv import load_dotenv

load_dotenv()

bot_token: Final = os.environ.get("BOT_TOKEN")

redis_host: Final = os.environ.get("REDIS_HOST")
redis_port: Final = int(os.environ.get("REDIS_PORT"))
redis_db: Final = os.environ.get("REDIS_DB")
redis_password: Final = os.environ.get("REDIS_PASSWORD")


def get_bot_token():
    return bot_token


@dataclass(frozen=True)
class RedisConfig:
    host: str
    port: int
    db: str | int
    password: str


redis_config = RedisConfig(
    host=redis_host, port=redis_port, db=redis_db, password=redis_password
)


def get_redis_config() -> RedisConfig:
    return redis_config
