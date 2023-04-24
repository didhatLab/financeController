import aiohttp
from string import Template

from src.models.currency import CurrencyRate


class OpenExchangerApi:
    api = "https://openexchangerates.org/api/"
    latest = Template(f"{api}latest.json?app_id=$app_id")

    def __init__(self, session: aiohttp.ClientSession, client_id: str):
        self._session = session
        self._client_id = client_id

    async def latest_currency(self) -> CurrencyRate:
        async with self._session.get(self.latest.substitute(app_id=self._client_id)) as resp:
            json = await resp.json()

        return CurrencyRate.from_dict(json)



