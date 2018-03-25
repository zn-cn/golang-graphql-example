FROM golang:1.10
WORKDIR /go/src/app
COPY ./src src
ENV GOPATH "/go/src/app"
RUN go get -d -v ./... && \
    go build -o main src/main.go && \
    rm -rf src
ENV APP_ADDR ":1323"
ENV GO_ENV "dev"
EXPOSE 1323

ENTRYPOINT main
