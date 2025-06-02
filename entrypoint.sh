#!/bin/bash


while ! nc $DB_HOST $DB_PORT; do
    echo "Postgres is unavailable - sleeping"
    sleep 2
done

exec "$@"