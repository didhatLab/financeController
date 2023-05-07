from aiogram.dispatcher.router import Router
from aiogram.filters import StateFilter
from aiogram import types
import redis.asyncio as aioredis
from aiogram.fsm.context import FSMContext

from src.routes.states import SettingState, UnlinkState
from src.routes.keyboard import get_setting_keyboard, get_start_keyboard
from src.services.notification import turn_on_notify, turn_off_notify


def notification_setting_route(redis: aioredis.Redis):
    route = Router()

    @route.message(StateFilter(SettingState.menu_setting))
    async def setting_page(message: types.Message, state: FSMContext):
        command = message.text

        if command == "Turn on notify":
            await turn_on_notify(redis, message.chat.id)
            await message.answer(
                "notifications activated!", reply_markup=get_setting_keyboard()
            )
        elif command == "Turn off notify":
            await turn_off_notify(redis, message.chat.id)
            await message.answer("notification disabled!")
        elif command == "Unlink Account":
            await state.set_state(UnlinkState.unlink)
            await redis.delete(str(message.chat.id))
            await message.answer("account unlinked", reply_markup=get_start_keyboard())
        else:
            await message.answer(
                "unknown command, chose command from keyboard",
                reply_markup=get_setting_keyboard(),
            )
