include .env

stop_containers:
	@echo "Stopping other docker containers"
	@if not "$(shell docker ps -q)" == "" ( \
		echo "found and stopped containers"; \
		docker stop $(shell docker ps -q); \
	) else ( \
		echo "no containers running..."; \
	)

create_container:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:16.3

create_db:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

start_container:
	docker start ${DB_DOCKER_CONTAINER}

create_migrations:
	goose create init sql

up_migrations: 
	goose -dir migrations postgres "postgres://root:secret@localhost:5432/rest_db?sslmode=disable" up

down_migrations:
	goose -dir migrations postgres "postgres://root:secret@localhost:5432/rest_db?sslmode=disable" down

build:
	@if exist ${BINARY} ( \
		del ${BINARY} /Q >nul 2>&1 && echo "Deleted ${BINARY}" \
	) else ( \
		echo "${BINARY} not found" \
	)
	@echo "Building Binary..."
	go build -o ${BINARY} ./cmd/server/.

run: build
	./${BINARY} 
#	@echo "api started..."

# 	can do all this manually specialy the build and run but just for QOL and learning how to make a makefile