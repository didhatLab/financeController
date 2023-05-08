VENV=venv
PYTHON=$(VENV)/bin/python3
PIP=$(VENV)/bin/pip

SHELL := /bin/bash

migrate:
	alembic upgrade head

venv:
	mkdir venv
	python3.11 -m venv venv
	$(PIP) install -r requirements.txt

prepare:
	docker-compose up -d --build postgres
	echo "===========wait postgres startup=============="
	sleep 8
	alembic upgrade head

start:
	docker-compose up -d --build mongodb redis telegram-notifier resolver finances auth back2front currency currency_updatrer

stop:
	docker-compose down --remove-orphans


