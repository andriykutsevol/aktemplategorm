#!/usr/bin/env bash
source ${1}


DB_CONTAINER=$(docker ps --filter "name=db_service" --format "{{.Names}}" | head -n 1)
if [ -z "$DB_CONTAINER" ]; then
    echo "Error: No running container found for db_service!"
    exit 1
fi
echo "Waiting for $DB_CONTAINER to become healthy..."
while [ "$(docker inspect --format='{{.State.Health.Status}}' $DB_CONTAINER 2>/dev/null)" != "healthy" ]; do
    echo "Waiting for database..."
    sleep 3
done
echo "Database is healthy, running migrations..."


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
