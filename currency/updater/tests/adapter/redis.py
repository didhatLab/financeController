import pytest
import redis.asyncio as aioredis


@pytest.fixture()
async def redis():
    _redis = aioredis.Redis()
    yield _redis
    await _redis.close()
