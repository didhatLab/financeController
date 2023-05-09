# Auth сервис

Сервис необходим для регистрации новых пользователей и для раздачи токенов уже существующим.
Для токенов используется механизм JWT токенов, что позволяет выдавать токены на ограниченное время, но 
сейчас они выдаются бессрочными для простоты реализации. 

Переменные окружения:

    POSTGRES_URI: uri для подключения к postgres
    AUTH_STARTUP_PORT: port на котором запустится сервис
    SECRET_KEY: ключ, с которым создаются токены и хэши паролей

Схема базы данных:

    user_id: int
    username: str
    pass_hash: str

API Entrypoints:

    POST /register { username: str, password: str }
    POST /token { username: str, password: str } - 200: {Token: str } or 400





