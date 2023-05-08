import os
from typing import Final

from dotenv import load_dotenv

load_dotenv()

secret_key: Final = os.environ.get("SECRET_KEY")

spend_url: Final = os.environ.get("SPENDING_URL")
currency_url: Final = os.environ.get("CURRENCY_URL")


def get_secret_key():
    return secret_key


def get_spend_url():
    return spend_url


def get_currency_url():
    return currency_url
