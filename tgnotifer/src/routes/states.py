from aiogram.fsm.state import State, StatesGroup


class LoginAccount(State):
    input_username = State()
    input_password = State()


class SettingState(State):
    menu_setting = State()
