install-dependencies:
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go mod tidy

build:
	@go build -o ./bin/main ./main.go 

run:
	./bin/main

sql-gen:
	@sqlc generate

migrate:
	@goose up

clean:
	@rm -rf ./bin/*
