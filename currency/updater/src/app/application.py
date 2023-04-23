import aiohttp
import redis.asyncio as aioredis
from motor.motor_asyncio import AsyncIOMotorDatabase

from src.api.openexchanger import OpenExchangerApi
from src.adapter.history_currency import CurrencyHistoryRepositoryImp
from src.adapter.current_currency import CurrentCurrencyRatesRepositoryImp
from src.services.currency_updater import CurrencyUpdater


class UpdaterCurrencyApplication:
    def __init__(
        self,
        mongo: AsyncIOMotorDatabase,
        redis: aioredis.Redis,
        session: aiohttp.ClientSession,
    ):
        self._mongo = mongo
        self._redis = redis
        self._api_session = session

    def build(self):
        client_id = "FIXME to env"

        currency_api = OpenExchangerApi(session=self._api_session, client_id=client_id)

        history_currency_repo = CurrencyHistoryRepositoryImp(
            self._mongo, "currency_history"
        )
        current_current_repo = CurrentCurrencyRatesRepositoryImp(
            self._redis, "current_currency_rate"
        )

        currency_updater_service = CurrencyUpdater(
            currency_api, current_current_repo, history_currency_repo
        )

        return currency_updater_service


