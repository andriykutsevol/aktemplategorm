#!/usr/bin/env bash
source ${1}


echo "[ ! ] Migrating database baseline"
echo "NETWORK_NAME: ${NETWORK_NAME}"
docker run \
    --network ${2} \
    -v ${MIGRATIONS}:/flyway/sql \
    --rm \
    flyway/flyway \
    -url="jdbc:mysql://${3}:${4}?sslMode=disable&allowPublicKeyRetrieval=true" \
    -user=${MYSQL_USER} \
    -password=${MYSQL_PASSWORD} \
    -schemas=${MYSQL_DATABASE} \
    -baselineVersion=${5} \
    baseline