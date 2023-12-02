.PHONY: default run build clean

default: run

run: 
	@go run cmd/main.go

build: 
	@go build -o cmd/out.o cmd/main.go

clean: 
	@rm -rf cmd/*.o