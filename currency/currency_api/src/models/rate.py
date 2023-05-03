from pydantic import BaseModel


class CurrencyRate(BaseModel):
    base: str
    timestamp: int
    rates: dict[str, float | int]


class RateForStats(BaseModel):
    Usd: float
    Eur: float