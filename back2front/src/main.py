import asyncio

import aiohttp
from fastapi import FastAPI, Depends
from fastapi.middleware.cors import CORSMiddleware

from src.auth import auth_header_token
from src.models.infalte import InflateUIData
from src.models.spend import Spending, CurrencyRate
from src.config import get_spend_url, get_currency_url


app = FastAPI()


async def aiohttp_session() -> aiohttp.ClientSession:
    session = aiohttp.ClientSession()
    yield session
    await session.close()


@app.get("/inflate_ui")
async def inflate_ui_for_app(
    info: int = Depends(auth_header_token),
    session: aiohttp.ClientSession = Depends(aiohttp_session),
) -> InflateUIData:
    user_id, token = info

    auth_header = {"Auth-Token": token}

    async with asyncio.TaskGroup() as tg:
        spends = tg.create_task(get_spends(session, auth_header))
        currency_rate = tg.create_task(get_currency_rate(session))

    return InflateUIData(spends=spends.result(), currency_rate=currency_rate.result())


async def get_spends(session: aiohttp.ClientSession, headers: dict) -> list[Spending]:
    async with session.post(f"{get_spend_url()}/spending/get", headers=headers) as r:
        spends = await r.json()

    return [Spending(**sp) for sp in spends]


async def get_currency_rate(session: aiohttp.ClientSession) -> CurrencyRate:
    async with session.get(f"{get_currency_url()}/current_currency_rate") as r:
        rate = await r.json()

    return CurrencyRate(**rate)


app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, port=4003)
