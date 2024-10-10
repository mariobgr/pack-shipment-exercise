FROM golang:latest

EXPOSE 8080

COPY . /go/src/app

WORKDIR /go/src/app

CMD ["sh", "-c", "go run /go/src/app/cmd/server/main.go"]
