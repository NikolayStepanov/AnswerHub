DOCKER_YML=${CURDIR}/docker-compose.yml
ENV_NAME=answer-hub

.PHONY: build-all
build-all:
	cd answer && GOOS=linux GOARCH=amd64 make build

.PHONY: run-all
run-all: build-all
	docker-compose -p ${ENV_NAME} up --force-recreate --build

.PHONY: compose-up
compose-up:
	docker-compose -p ${ENV_NAME} -f ${DOCKER_YML} up -d

.PHONY: compose-down
compose-down:
	docker-compose -p ${ENV_NAME} -f ${DOCKER_YML} stop

.PHONY: compose-rm
compose-rm:
	docker-compose -p ${ENV_NAME} -f ${DOCKER_YML} rm -fvs

.PHONY: compose-rs
compose-rs:
	make compose-rm
	make compose-up