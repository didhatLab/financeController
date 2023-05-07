from aiogram.fsm.state import State, StatesGroup


class LoginAccount(StatesGroup):
    input_username = State()
    input_password = State()


class SettingState(StatesGroup):
    menu_setting = State()


class UnlinkState(StatesGroup):
    unlink = State()
