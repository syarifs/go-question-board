run:
	@go run main.go

test:
	@go test -cover -coverpkg=./... ./...

test_html:
	@mkdir /tmp/go
	@go test -coverprofile=/tmp/go/test.out -coverpkg=./... ./...
	@go tool cover -html=/tmp/go/test.out
	@rm -rf /tmp/go

build:
	@go build -o bin/question-board main.go

docs:
	@swag init

tidy:
	@go mod tidy
