from pydantic import BaseModel

from src.models.spend import CurrencyRate, Spending


class InflateUIData(BaseModel):
    spends: list[Spending]
    currency_rate: CurrencyRate
