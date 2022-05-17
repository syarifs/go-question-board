FROM golang:1.18 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app main.go

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /
COPY --from=build-env /go/src/app/config.yaml /
EXPOSE 8080
CMD ["/app"]
