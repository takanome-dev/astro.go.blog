BINARY_NAME=blog-api
.DEFAULT_GOAL := run

migrate:
	cd sql/schema && ~/go/bin/goose postgres $(DB_URL) up

install:
	go mod tidy && go mod vendor

build:
	go build -o bin/$(BINARY_NAME) main.go

run: build
	bin/$(BINARY_NAME)

clean: 
	rm -r bin