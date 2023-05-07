from aiogram import types


settings_menu = [
    [
        types.KeyboardButton(text="Turn on notify"),
        types.KeyboardButton(text="Turn off notify"),
    ],
    [types.KeyboardButton(text="Unlink Account")],
]

setting_keyboard = types.ReplyKeyboardMarkup(
    keyboard=settings_menu, resize_keyboard=True, input_field_placeholder="Settings"
)


def get_setting_keyboard():
    return setting_keyboard
