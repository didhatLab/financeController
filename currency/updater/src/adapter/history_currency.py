import datetime

from motor.motor_asyncio import AsyncIOMotorDatabase

from src.models.currency import CurrencyRate


class CurrencyHistoryRepositoryImp:
    def __init__(self, mongo: AsyncIOMotorDatabase, collection_name: str):
        self._mongo = mongo
        self._collection_name = collection_name

    async def get_rates_for_time(self, time: datetime.datetime) -> CurrencyRate:
        pass

    @property
    def _collection(self):
        return self._mongo.get_collection(self._collection_name)

