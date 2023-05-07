import os
from typing import Final
from dataclasses import dataclass

from dotenv import load_dotenv

load_dotenv()

mongo_uri: Final = os.environ.get("MONGO_URI")
mongo_db: Final = os.environ.get("MONGO_DATABASE")

redis_host: Final = os.environ.get("REDIS_HOST")
redis_db: Final = os.environ.get("REDIS_DB")
redis_port: Final = int(os.environ.get("REDIS_PORT"))
redis_password: Final = os.environ.get("REDIS_PASSWORD")


@dataclass(frozen=True)
class RedisConfig:
    host: str
    port: int
    db: str | int
    password: str | None


@dataclass(frozen=True)
class MongoConfig:
    uri: str
    db: str


redis_config = RedisConfig(
    host=redis_host, port=redis_port, db=redis_db, password=redis_password
)

mongo_config = MongoConfig(uri=mongo_uri, db=mongo_db)


def get_redis_config():
    return redis_config


def get_mongo_config():
    return mongo_config
