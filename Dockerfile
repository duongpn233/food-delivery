FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin/main ./cmd

EXPOSE 8080

ENV CONFIG_PATH="config/config.toml"

CMD ["./bin/main", "-c", "${CONFIG_PATH}"]