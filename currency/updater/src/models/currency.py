import attr
import cattr
import enum

from typing import AnyStr, Union


class Currency(enum.StrEnum):
    usd = "USD"
    rub = "RUB"
    eur = "EUR"


@attr.dataclass
class CurrencyRate:
    timestamp: int
    base: Currency
    rates: dict[Union[str, Currency], Union[float, int]]

    @classmethod
    def from_dict(cls, raw: dict) -> "CurrencyRate":
        return converter.structure(raw, cls)

    def to_dict(self):
        return converter.unstructure(self)


converter = cattr.Converter()

converter.register_structure_hook(Union[str, Currency], lambda value, _: value)
converter.register_structure_hook(Union[float, int], lambda value, _: value)
