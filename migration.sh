#!/bin/sh

PG_USERNAME="postgres"
PG_PASSWORD="qwerty"
PG_HOST="localhost"
PG_PORT="5432"
PG_DBNAME="hardtry-db"
PG_SSLMODE="disable"

MIGRATIONS_DIR="./migrations"

export PG_DSN="host=${PG_HOST} port=${PG_PORT} dbname=${PG_DBNAME} user=${PG_USERNAME} password=${PG_PASSWORD} ${PG_SSLMODE}=disable"
goose -dir ${MIGRATIONS_DIR} postgres "${PG_DSN}" up -v
