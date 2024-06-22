FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /go-redis

EXPOSE 8080

CMD ["/go-redis"]

