run:
	go run main.go

test:
	go test -cover -coverpkg=./... ./...

build:
	go build -o bin/question-board main.go

docs:
	swag init

tidy:
	go mod tidy
