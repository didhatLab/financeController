import os
from typing import Final
from dataclasses import dataclass

from dotenv import load_dotenv

load_dotenv()

redis_host: Final = os.environ.get("REDIS_HOST")
redis_db: Final = os.environ.get("REDIS_DB")
redis_port: Final = int(os.environ.get("REDIS_PORT"))
redis_password: Final = os.environ.get("REDIS_PASSWORD")

startup_port: Final = int(os.environ.get("STARTUP_PORT"))


@dataclass(frozen=True)
class RedisConfig:
    host: str
    port: int
    db: str | int
    password: str | None


redis_config = RedisConfig(
    host=redis_host, port=redis_port, db=redis_db, password=redis_password
)


def get_redis_config():
    return redis_config


def get_startup_config():
    return startup_port
