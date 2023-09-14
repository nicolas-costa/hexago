FROM golang as base

WORKDIR /app

COPY . .

RUN go mod download

FROM base as dev

RUN go install github.com/githubnemo/CompileDaemon@latest && \
  go install github.com/joho/godotenv/cmd/godotenv@latest

RUN go build -o ./cmd/hexago/hexago ./main.go

ENTRYPOINT [ "./cmd/hexago/hexago" ]