import asyncio
from typing import List

import aiohttp
from fastapi import FastAPI, Depends
from pydantic import parse_obj_as

from src.auth import auth_header_token
from src.models.infalte import InflateUIData
from src.models.spend import Spending, CurrencyRate


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
    async with session.post("http://127.0.0.1:4000/spending/get", headers=headers) as r:
        spends = await r.json()

    return [Spending(**sp) for sp in spends]


async def get_currency_rate(session: aiohttp.ClientSession) -> CurrencyRate:
    async with session.get("http://127.0.0.1:4002/current_currency_rate") as r:
        rate = await r.json()

    return CurrencyRate(**rate)


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, port=4003)
