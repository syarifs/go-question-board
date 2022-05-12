run:
	@clear
	@go run main.go

test:
	@clear
	@go test -cover -coverpkg=./... ./...

test_html:
	@clear
	@mkdir /tmp/go
	@go test -coverprofile=/tmp/go/test.out -coverpkg=./... ./...
	@go tool cover -html=/tmp/go/test.out
	@rm -rf /tmp/go

build:
	@clear
	@go build -o bin/question-board main.go

docs:
	@clear
	@swag init

tidy:
	@go mod tidy
