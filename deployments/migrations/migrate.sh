#!/usr/bin/env bash
source ${1}

echo "[ ! ] Migrating database"
docker run \
    --network ${2} \
    -v ${MIGRATIONS}:/flyway/sql \
    --rm \
    flyway/flyway:latest \
    -url="jdbc:mysql://${3}:${4}?sslMode=disable&allowPublicKeyRetrieval=true" \
    -user=${MYSQL_USER} \
    -password=${MYSQL_PASSWORD} \
    -schemas=${MYSQL_DATABASE} \
    migrate
