import pytest
import json

from src.adapter.current_currency import CurrentCurrencyRatesRepositoryImp
from src.models.currency import CurrencyRate, Currency


some_time = 1682253779
currency_test_key = "test_currency"


@pytest.fixture()
def currency_rate():
    return CurrencyRate(
        timestamp=some_time,
        base=Currency.usd,
        rates={Currency.usd: 100, Currency.rub: 81.2},
    )


@pytest.fixture()
def current_currency_repository(redis) -> CurrentCurrencyRatesRepositoryImp:
    repo = CurrentCurrencyRatesRepositoryImp(redis, currency_test_key)

    return repo


async def test_set_new_current_currency_rate(
    current_currency_repository, currency_rate: CurrencyRate, redis
):
    await current_currency_repository.save_new_rate(currency_rate)

    json_str = await redis.get(currency_test_key)

    raw_dict = json.loads(json_str)

    assert CurrencyRate.from_dict(raw_dict) == currency_rate
