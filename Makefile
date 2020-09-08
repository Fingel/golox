build:
	go build -o bin/lox lox.go
run:
	go run lox.go
format:
	gofmt -s -w .
