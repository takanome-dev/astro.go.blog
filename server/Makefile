BINARY_NAME=blog-api
.DEFAULT_GOAL := run

migrate:
	cd sql/schema && goose postgres $(DB_URL) up

install:
	go mod tidy

generate:
	sqlc generate

build:
	go build -o bin/$(BINARY_NAME) main.go

run: build
	bin/$(BINARY_NAME)

lint:
	golangci-lint run

clean: 
	rm -r bin