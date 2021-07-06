build:
	@go build -v -o main main.go 

run:
	make build
	@./main