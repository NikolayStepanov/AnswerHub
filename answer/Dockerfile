FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o bin/answer-server ./cmd/app/

FROM alpine

COPY --from=builder /app/bin/answer-server /app/answer-server
COPY --from=builder /app/configs /app/configs/

ENV CONFIG_DIR=/app/configs

CMD [ "/app/answer-server" ]
