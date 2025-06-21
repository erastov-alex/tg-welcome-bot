FROM golang:1.23.10-bullseye

WORKDIR /app

COPY go.mod go.sum ./
COPY vendor/ ./vendor/
COPY . .

RUN go build -mod=vendor -o tg-welcome-bot ./cmd/bot

CMD ["./tg-welcome-bot"]
