.SILENT: ;				 # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

SHELL          := $(shell which bash)
ROOT_DIR       := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BACKEND_DIR	   := ${ROOT_DIR}/cmd/mus
PROJECT_NAME   := mus
VERSION        := 1.0.0
MIGRATIONS	   := $(ROOT_DIR)/deployments/migrations
DEPLOYMENTS    := $(ROOT_DIR)/deployments/compose

#	The --progress=plain is mandatory if you want to look at RUN ls /opt/app command
# 	You could also disable buildkit (not recommended)
#	export DOCKER_BUILDKIT=0
#   If you want to keep buildkit and make it print commands output, then:
BUILDKIT_PROGRESS=plain
#BUILDKIT_PROGRESS=auto	# (default): Allows BuildKit to choose the most suitable output format.
#BUILDKIT_PROGRESS=tty	# Uses a fancy progress display that groups and summarizes each stage of the build. 
#						This is the default when BuildKit is enabled and the terminal supports TTY.

export BUILDKIT_PROGRESS
#	or use --progress=plain with the "docker build" or "docker-compose ... build ..."

ENV ?= dev
ifeq ($(ENV), dev)
ENVFILE = .env_dev
endif
ifeq ($(ENV), test)
ENVFILE = .env_test
endif
include $(ENVFILE)

.DEFAULT: help
.PHONY: help
help:
	echo "=============="
	echo "Show Help Menu"
	echo "make clear"
	echo "make run"
	echo "make migrate_baseline-onlocalhost"
	echo "make migrate-onnetwork"
	echo "=============="


#============================================================================
# DB related.
# Usecases:

# 1) run_db-onnetwork
# 2) migrate_baseline-onnetwork
# 3) migrate-onnetwork

# 1) run_db-onnetwork
# 2) migrate_baseline-onlocalhost
# 3) migrate-onlocalhost
#============================================================================

.PHONY: run_db-onnetwork
run_db-onnetwork:
	/bin/bash $(DEPLOYMENTS)/create_network.sh
#	This command creates simple volume because we define it on our .yml
#	And when we delete a container, the volume remains, and when we create and launch a container, 
#	this volume is mounted (because it is specified in .yml) and if there is already data there, it is saved.
	docker-compose --project-name $(PROJECT_NAME) -f $(DEPLOYMENTS)/docker-compose.yml --profile db up -d

.PHONY: clear_db
clear_db-onnetwork:
	docker-compose --project-name $(PROJECT_NAME) -f $(DEPLOYMENTS)/docker-compose.yml --profile db down
	docker volume rm $(PROJECT_NAME)_db
	/bin/bash $(DEPLOYMENTS)/remove_network.sh

# You need to wait like 10 seconds after run_db before this.
.PHONY: migrate_baseline-onnetwork
migrate_baseline-onnetwork:
	/bin/bash $(MIGRATIONS)/baseline.sh $(ROOT_DIR)/$(ENVFILE) $(NETWORK_NAME) $(MYSQL_HOST_ONNETWORK) $(MYSQL_PORT_ONNETWORK) $(FLYWAY_BASELINE)

.PHONY: migrate_baseline-onlocalhost
migrate_baseline-onlocalhost:
	/bin/bash $(MIGRATIONS)/baseline.sh $(ROOT_DIR)/$(ENVFILE) host 0.0.0.0 $(MYSQL_PORT_ONHOST) $(FLYWAY_BASELINE)


.PHONY: migrate-onnetwork
migrate-onnetwork:
	/bin/bash $(MIGRATIONS)/migrate.sh $(ROOT_DIR)/$(ENVFILE) $(NETWORK_NAME) $(MYSQL_HOST_ONNETWORK) $(MYSQL_PORT_ONNETWORK)


.PHONY: migrate-onlocalhost
migrate-onlocalhost:
	/bin/bash $(MIGRATIONS)/migrate.sh $(ROOT_DIR)/$(ENVFILE) host 0.0.0.0 $(MYSQL_PORT_ONHOST)




#============================================================================
# Backend related.
#============================================================================

# Use cases:
# RUN without docker
# 1) run_go-onlocalhost

# Copy project to container, build it, and run.
# It is fast, it do not redownload required dependencies
# but it'll copy whole project to container even if one file changed.
# (it is how docker layer caching works (COPY it's a layer))
# 1) build_go_dev
# 2) up_go_dev-onnetwork

# The same as above but with single command
# 	You make changes in your code, then run this.
# 	It is fast, it do not redownload required dependencies.
# 	But it copy whole project to container every time.
# 	To avoid this we need to implement bind mount (for development)
# 1) rebuid_go_dev-onnetwork


.PHONY: run_go-onlocalhost
run_go-onlocalhost:
	export DATABASE_DSN=$(DATABASE_DSN_ONHOST)
	export APP_PORT=$(APP_PORT_ONLOCALHOST)
	go run $(BACKEND_DIR)/.


.PHONY: build_go_dev
build_go_dev:
# 	The correct way to specify a custom Compose file is to put it BEFORE the subcommand, not as part of the build command:
	docker-compose -f $(DEPLOYMENTS)/docker-compose.yml build backend_dev

.PHONY: up_go_dev-onnetwork
up_go_dev-onnetwork:
	export DATABASE_DSN=$(DATABASE_DSN_ONNETWORK)
	export APP_PORT=$(APP_PORT_ONNETWORK)
	docker-compose --project-name $(PROJECT_NAME) -f $(DEPLOYMENTS)/docker-compose.yml --profile backend_dev up


# You make changes in your code, then run this.
# It is fast, it do not redownload required dependencies.
# But it copy whole project to container every time.
# To avoid this we need to implement bind mount (for development)
.PHONY: rebuid_go_dev-onnetwork
rebuid_go_dev-onnetwork: build_go_dev up_go_dev-onnetwork
	echo "All targets rebuid_go_dev-onnetwork executed!"


#---------------------------------------------------------------------------
# Use case:
# It's when you want to run container with -it and /bin/bash.
# The docker-compose up does not support -it
# But "run" will create new container every time, and it does not understand --project-name
# And that DEBUG_MODE indeed override the CMD in dockerfile.
# We need this DEBUG_MODE if our container quit quickly (in this case we change CMD to /bin/bash)

# Use case:
# 1) build_go_dev
# 2) debug_go_dev-onnetwork
# 3) it will show container's bash prompt, and you can go run ./cmd/mus
# and it'll redownload dependencies every time you do make debug_go_dev-onnetwork
# because it'll create new container everytime. But why it do not cache it during image build - I dunno.

.PHONY: debug_go_dev-onnetwork
debug_go_dev-onnetwork:
	export DATABASE_DSN=$(DATABASE_DSN_ONNETWORK)
	export APP_PORT=$(APP_PORT_ONNETWORK)
	export DEBUG_MODE="/bin/bash"
	docker-compose -f $(DEPLOYMENTS)/docker-compose.yml run -it --rm backend_dev


#---------------------------------------------------------------------------
# Use case:
# We want to connect to already running backend_dev service
# 1) make rebuid_go_dev-onnetwork
# 2) make bash_backend_dev

# we can use this only if the up_go_dev-onnetwork has the export DEBUG_MODE="/bin/bash" uncommented.
# Why? Because then I'll wait, if your go app do not quit quickly, you do not need thet export.
.PHONY: bash_backend_dev
bash_backend_dev:
	export DATABASE_DSN=$(DATABASE_DSN_ONNETWORK)
	export APP_PORT=$(APP_PORT_ONNETWORK)
	docker-compose --project-name $(PROJECT_NAME) -f $(DEPLOYMENTS)/docker-compose.yml exec -it backend_dev /bin/bash





#============================================================================
# Backend related with volume
#============================================================================
# We want to run our code within docker's network but we also want to
# change an rebuild our code quickly - so we use volume.
# If we'll still use
# COPY go.mod go.sum ./
# and
# RUN --mount=type=cache,target=/go/pkg/mod \
#     --mount=type=cache,target=/root/.cache/go-build \
#     go mod tidy 
# in the docker file to use cache during building
# then that cache will NOT be included in resulting image
# that's why when you run new container and then run go run ./cmd/mus
# it'll redownload all dependencies (but it's only when you recreate container)
# that's why I removed it from docker image.
# The --mount=type=cache option is specific to the build process and is not applicable 
# when running a container from the built image.

# Use case:
# 1) build_go_dev_vol
# 2) up_go_dev_vol-onnetwork
# 3) bash_backend_dev_vol		# we need this step because the "docker-compose up" does not support -it

.PHONY: build_go_dev_vol
build_go_dev_vol:
#   We do not need to do this for build step, they're visiable in .yml file without this export.
#   We do this here to make it clear which variables are actually needed.
	export ROOT_DIR
	export DEPLOYMENTS
# 	The correct way to specify a custom Compose file is to put it BEFORE the subcommand, not as part of the build command:	
	docker-compose -f $(DEPLOYMENTS)/docker-compose.yml build --no-cache backend_dev_vol

# Use case:
# If run this when we have stopped container (created with build_go_dev_vol) it does not redownload dependencies.
.PHONY: up_go_dev_vol-onnetwork
up_go_dev_vol-onnetwork:
#	We do not need export all (except of DATABASE_DSN and APP_PORT, they're not in the .env_dev)
#	We export it here to make it clear which variables are actually needed 
	export DATABASE_DSN=$(DATABASE_DSN_ONNETWORK)
	export MYSQL_MAX_OPENCONNS
	export MYSQL_MAX_IDLECONS
	export APP_PORT=$(APP_PORT_ONNETWORK)
	export APP_PORT_ONLOCALHOST
	export APP_PORT_ONNETWORK
	export BACKEND_HOST_ONNETWORK
	export ROOT_DIR
	docker-compose --project-name $(PROJECT_NAME) -f $(DEPLOYMENTS)/docker-compose.yml --profile dev_vol up


.PHONY: create_go_dev_vol-onnetwork
create_go_dev_vol-onnetwork:
#	We do not need export all (except of DATABASE_DSN and APP_PORT, they're not in the .env_dev)
#	We export it here to make it clear which variables are actually needed 
	export DATABASE_DSN=$(DATABASE_DSN_ONNETWORK)
	export MYSQL_MAX_OPENCONNS
	export MYSQL_MAX_IDLECONS
	export APP_PORT=$(APP_PORT_ONNETWORK)
	export APP_PORT_ONLOCALHOST
	export APP_PORT_ONNETWORK
	export BACKEND_HOST_ONNETWORK
	export ROOT_DIR
	docker-compose --project-name $(PROJECT_NAME) -f $(DEPLOYMENTS)/docker-compose.yml --profile dev_vol create


# Use case:
# 1) rebuid_go_dev_vol-onnetwork #  Just to combine the build_go_dev_vol and up_go_dev_vol-onnetwork

.PHONY: rebuid_go_dev_vol-onnetwork
rebuid_go_dev_vol-onnetwork: build_go_dev_vol up_go_dev_vol-onnetwork
	echo "All targets executed!"

# We need to stop container from docker-compose-desktop at leas once.
# Then I'll save it's state and it will not rebuil whole app each time you run it.
# It also will preserve commands history.
.PHONY: bash_backend_dev_vol
bash_backend_dev_vol:
	docker-compose --project-name $(PROJECT_NAME) -f $(DEPLOYMENTS)/docker-compose.yml exec -it backend_dev_vol /bin/bash


# Use case:

# This "run" creates new container every time (it does not understand PROJECT_NAME)
#  If you need persistent behavior and avoid the random suffix, 
#  switch to docker-compose up with proper service definitions in docker-compose.yml
.PHONY: debug_go_dev_vol-onnetwork
debug_go_dev_vol-onnetwork:
	export DATABASE_DSN=$(DATABASE_DSN_ONNETWORK)
	export APP_PORT=$(APP_PORT_ONNETWORK)
	export COMMAND="/bin/bash"
	docker-compose --project-name $(PROJECT_NAME) -f $(DEPLOYMENTS)/docker-compose.yml run -it --rm backend_dev_vol




.PHONY: backend_dev_vol_fromscratch
backend_dev_vol_fromscratch: run_db-onnetwork migrate_baseline-onnetwork migrate-onnetwork build_go_dev_vol create_go_dev_vol-onnetwork
	echo "All targets executed!"
# Now you can run mus_backend_dev_vol from docker compose (or by up_go_dev_vol-onnetwork )
# Then you do bash_backend_dev_vol and you're in backend's console:
# go run ./cmd/mus

#============================================================================
#============================================================================

.PHONY: redocly
redocly:
	docker run --rm -v ${ROOT_DIR}/api/$(API_VERSION):/spec redocly/cli bundle /spec/openapi.yaml --output /spec/bundled.yaml
	docker run --rm -v ${ROOT_DIR}/api/$(API_VERSION):/spec redocly/cli lint /spec/bundled.yaml




#============================================================================
#============================================================================



# make generate_mocks_pkg
MOCKS		   := $(ROOT_DIR)/mocks
MOCKPKG ?= './'

all:
    @echo "MOCKPG is set to: $(MOCKPKG)"


# Can I tell to the golang mockery to generate packages according to folder names not "package mocks"
# YES: Using --with-package
# mockery --dir ./internal/adapters/driving/restapi/router --output ./mocks/internal/adapters/driving/restapi/router --with-package


.PHONY: generate_mocks_all
generate_mocks_all:
	mockery --all --keeptree --output $(MOCKS)

.PHONY: clear_mocks_all
clear_mocks_all:
	rm -rf $(MOCKS)

# $ make generate_mocks_pkg MOCKS=./mocks MOCKPKG=./mockeryx
# 	When MOCKPKG is passed via the command line, make treats it as a defined variable, 
#	but ifndef only checks if the variable is entirely undefined, not if it's empty.
# 	Instead of checking if MOCKPKG is undefined, check if it's empty.


# Overriding via Command Line

.PHONY: generate_mocks_pkg
generate_mocks_pkg:
	echo "MOCKPKG is set to: $(MOCKPKG)"
	if [ -z "$(MOCKPKG)" ]; then \
		echo "MOCKPKG is not set. Please set MOCKPKG to the directory containing the package to generate mocks for."; \
		exit 1; \
	fi
	mockery --all --keeptree --output $(MOCKS) --dir $(MOCKPKG)



# In Makefile, a single $ is used for variable substitution.
# In a Makefile, $(...) is interpreted as a Make variable substitution, not a shell command.
# As a result, $(go list ./... | grep -v ./integration) is not executed as expected.
# To pass a literal $ to the shell (so it executes $(...) properly), you must escape it with another $, making it $$().

.PHONY: run_unit-tests
run_unit-tests:
	go test $$(go list ./... | grep -v ./integration)