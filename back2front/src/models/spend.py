from pydantic import BaseModel


class Spending(BaseModel):
    Id: int
    Name: str
    Type: str
    Amount: int
    Currency: str

    class Config:
        frozen = True


class CurrencyRate(BaseModel):
    base: str
    timestamp: int
    rates: dict[str, float | int]

    class Config:
        frozen = True
