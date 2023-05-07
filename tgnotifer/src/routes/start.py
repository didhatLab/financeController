from aiogram.dispatcher.router import Router
from aiogram.filters import Command, StateFilter
from aiogram import types
from aiogram.fsm.context import FSMContext

from src.routes.keyboard import get_start_keyboard
from src.routes.states import UnlinkState, LoginAccount


def start_route():
    router = Router()

    @router.message(Command(commands=["start"]))
    async def start_menu(message: types.Message, state: FSMContext):
        await state.set_state(UnlinkState.unlink)
        await message.answer(
            "It is bot for notifications in financeController",
            reply_markup=get_start_keyboard(),
        )

    @router.message(StateFilter(UnlinkState.unlink))
    async def start_login_account(message: types.Message, state: FSMContext):
        await message.answer("Input login for your account")

        await state.set_state(LoginAccount.input_username)

    return router
