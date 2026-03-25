# Load .env file
include .env
export

#Variables
APP_NAME = easy-ride-api
MAIN = ./cmd

#Default target
.DEFAULT_GOAL	:= dev

##Run with Air (live reload)
dev:
	air -c air.conf

##Run with Go run (no live reload)
run:	
	go run $(MAIN)

##Build the application
build:
	go build -o bin/$(APP_NAME) $(MAIN)

##clean the build
clean:	
	rm -rf bin tmp

##start the application
start: build
	./bin/$(APP_NAME)

##Run tests
test:
	go test -v ./...


##Migration

DB_URL = "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

migrate-up:
	migrate -path migrations -database $(DB_URL) up

migrate-down:
	migrate -path migrations -database $(DB_URL) down

migrate-force:
	migrate -path migrations -database $(DB_URL) force $(version)