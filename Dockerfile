FROM golang:1.10
WORKDIR /go/src/app
COPY ./src src
ENV GOPATH "/go/src/app"
RUN go get -d -v ./...
ENV APP_ADDR ":1323"
ENV GO_ENV "dev"
EXPOSE 1323

ENTRYPOINT go run src/main.go
