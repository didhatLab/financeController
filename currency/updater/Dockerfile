FROM python:3.11

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

WORKDIR /app

COPY ./requirements.txt requirements.txt
COPY ./.env .env

RUN pip install -r requirements.txt

COPY ./src src

CMD ["python", "-m", "src.main"]
