#!/bin/bash

# Check if the network already exists
if ! docker network ls --format '{{.Name}}' | grep -q "^${NETWORK_NAME}\$"; then
  echo "Network ${NETWORK_NAME} does not exist. Creating it..."
  docker network create --subnet=172.18.1.0/24 --gateway=172.18.1.1 "${NETWORK_NAME}"
else
  echo "Network ${NETWORK_NAME} already exists."
fi


# flyway.url=
# jdbc:mysql://$(MYSQL_HOST):$(MYSQL_PORT)?user=$(MYSQL_USER)&password=$(MYSQL_PASSWORD)&sslMode=disable&useSSL=false&requireSSL=false

# .PHONY:
# migrate:
# 	echo "[ ! ] Migrating database"
# 	docker run \
# 	--network $(NETWORK_NAME) \
# 	-v $(MIGRATIONS):/flyway/sql \
# 	--rm \
# 	flyway/flyway \
# 	-url=jdbc:mysql://$(MYSQL_HOST):$(MYSQL_PORT)?sslMode=disable \
# 	-schemas=$(MYSQL_DATABASE) \
# 	-user=$(MYSQL_USER) \
# 	-password=$(MYSQL_PASSWORD) \
# 	migrate


# .PHONY:
# migrate:
# 	echo "[ ! ] Migrating database"
# 	docker run \
# 	--network $(NETWORK_NAME) \
# 	-v $(MIGRATIONS):/flyway/sql \
# 	--rm \
# 	flyway/flyway \
# 	-url=jdbc:mysql://$(MYSQL_HOST):$(MYSQL_PORT)?sslMode=disable \
# 	-schemas=$(MYSQL_DATABASE) \
# 	-user=$(MYSQL_USER) \
# 	-password=$(MYSQL_PASSWORD) \
# 	migrate



