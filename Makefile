build:
	@go build -o ./bin/main ./main.go 

run:
	./bin/main

clean:
	@rm -rf ./bin/*
