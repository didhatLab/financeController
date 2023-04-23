import datetime

import pytest
from motor.motor_asyncio import AsyncIOMotorDatabase

from src.models.currency import CurrencyRate, Currency
from src.adapter.history_currency import CurrencyHistoryRepositoryImp


collection_for_test = "test_currency"
just_time = 1682253779


@pytest.fixture()
def currency_rate() -> CurrencyRate:
    return CurrencyRate(
        timestamp=just_time,
        base=Currency.usd,
        rates={Currency.usd: 12, Currency.rub: 45, "FFF": 34},
    )


@pytest.fixture()
def currency_history_repository(mongo_currency) -> CurrencyHistoryRepositoryImp:
    repo = CurrencyHistoryRepositoryImp(mongo_currency, collection_for_test)
    return repo


async def test_save_rate_history(
    currency_history_repository,
    currency_rate: CurrencyRate,
    mongo_currency: AsyncIOMotorDatabase,
):
    await currency_history_repository.save_to_history_rate(currency_rate)

    check = await mongo_currency.get_collection(collection_for_test).find_one({})

    history_value = CurrencyRate.from_dict(check)

    assert history_value == currency_rate


async def test_get_rate_history(
    currency_history_repository, currency_rate: CurrencyRate
):
    await currency_history_repository.save_to_history_rate(currency_rate)

    currency_history = await currency_history_repository.get_rates_for_time(
        datetime.datetime.fromtimestamp(just_time)
    )

    assert currency_history == currency_rate
