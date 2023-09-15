FROM golang:1.21-alpine as base

WORKDIR /app

COPY . .

RUN go mod download

FROM base as dev

RUN go install github.com/githubnemo/CompileDaemon@latest && \
  go install github.com/joho/godotenv/cmd/godotenv@latest
