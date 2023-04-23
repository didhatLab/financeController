import pytest

from motor.motor_asyncio import AsyncIOMotorDatabase, AsyncIOMotorClient


@pytest.fixture()
async def mongo_currency() -> AsyncIOMotorDatabase:
    client = AsyncIOMotorClient("localhost:27017")
    db: AsyncIOMotorDatabase = client.get_database("currency_test")
    yield db
    await client.drop_database("currency_test")



