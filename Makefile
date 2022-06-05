run:
	@clear
	@go run main.go

test:
	@clear
	@rm -rf /tmp/go
	@mkdir /tmp/go
	@Env=testing go test -coverprofile=/tmp/go/test.out -cover -coverpkg=./... ./...
	@go tool cover -func /tmp/go/test.out

test_html:
	@clear
	@rm -rf /tmp/go
	@mkdir /tmp/go
	@Env=testing go test -coverprofile=/tmp/go/test.out -coverpkg=./... ./...
	@go tool cover -html=/tmp/go/test.out

build:
	@clear
	@go build -o bin/question-board main.go

build_docker:
	@clear
	@docker image build .

docs:
	@clear
	@swag init

tidy:
	@go mod tidy
