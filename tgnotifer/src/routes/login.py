from aiogram.dispatcher.router import Router
from aiogram.filters import StateFilter
from aiogram import types
import redis.asyncio as aioredis
from aiogram.fsm.context import FSMContext

from src.routes.states import LoginAccount, SettingState
from src.services.login import login, save_account, LoginException
from src.routes.keyboard import get_setting_keyboard


def login_menu_router(redis: aioredis):
    login_router = Router()

    @login_router.message(StateFilter(LoginAccount.input_username))
    async def get_username_for_login(message: types.Message, state: FSMContext):
        username = message.text

        await state.update_data(username=username)
        await state.set_state(LoginAccount.input_password)

        await message.answer("input password")

    @login_router.message(StateFilter(LoginAccount.input_password))
    async def try_login(message: types.Message, state: FSMContext):
        password = message.text

        data = await state.get_data()
        username = data.get("username")

        try:
            token, user_id = await login(username, password)
        except LoginException as e:
            await message.answer("Error in password or username, try again!")
            return

        await save_account(redis, message.chat.id, user_id, token)

        await message.answer(
            "Account linked, you can set notification settings now!",
            reply_markup=get_setting_keyboard(),
        )
        await state.set_state(SettingState.menu_setting)

    return login_router
