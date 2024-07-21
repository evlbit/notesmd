include .env

build:
	@go build -o bin/notesmd cmd/api/main.go

run: build
	@./bin/notesmd

mr-create:
	@read -p "Migration name: " MIGRATION_NAME; \
	goose -dir ${MIGRATION_DIR} create $${MIGRATION_NAME} sql

mr-up:
	@goose -dir ${MIGRATION_DIR} ${MIGRATION_DRIVER} ${MIGRATION_DBURL} up

mr-down:
	@goose -dir ${MIGRATION_DIR} ${MIGRATION_DRIVER} ${MIGRATION_DBURL} down

mr-reset:
	@goose -dir ${MIGRATION_DIR} ${MIGRATION_DRIVER} ${MIGRATION_DBURL} reset
