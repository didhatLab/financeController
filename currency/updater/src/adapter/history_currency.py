import datetime
from typing import cast, Optional, Any

from motor.motor_asyncio import (
    AsyncIOMotorDatabase,
    AsyncIOMotorCollection,
    AsyncIOMotorCursor,
)

from src.models.currency import CurrencyRate


class CurrencyHistoryRepositoryImp:
    def __init__(self, mongo: AsyncIOMotorDatabase, collection_name: str):
        self._mongo = mongo
        self._collection_name = collection_name

    async def get_rates_for_time(
        self, time: datetime.datetime
    ) -> Optional[CurrencyRate]:
        unix_time = int(time.timestamp())
        cursor: AsyncIOMotorCursor = self._collection.aggregate(
            [
                {
                    "$project": {
                        "diff": {"$abs": {"$subtract": [unix_time, "$timestamp"]}},
                        "doc": "$$ROOT",
                    }
                },
                {"$sort": {"diff": 1}},
                {"$limit": 1},
            ]
        )
        raw = await cursor.next()
        await cursor.close()
        if raw is None:
            return None

        return CurrencyRate.from_dict(raw["doc"])

    async def save_to_history_rate(self, rate: CurrencyRate) -> Any:
        await self._collection.insert_one(rate.to_dict())

    @property
    def _collection(self) -> AsyncIOMotorCollection:
        return cast(
            AsyncIOMotorCollection, self._mongo.get_collection(self._collection_name)
        )
