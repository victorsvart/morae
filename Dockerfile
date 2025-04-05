FROM golang:1.24-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN apk add --no-cache bash curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/

# Copy go.mod first
COPY go.mod ./

COPY go.sum* ./

COPY . .
RUN go mod download && \
    go mod tidy

# Build the application
RUN go build -o main ./cmd/api

EXPOSE 8080

CMD migrate -path=./cmd/migrate/migrations -database="$DSN" up && ./main
