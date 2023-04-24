from pydantic import BaseModel


class CurrencyRate(BaseModel):
    base: str
    timestamp: int
    rates: dict[str, float | int]
