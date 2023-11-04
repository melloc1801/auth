FROM golang:1.20-alpine3.17 AS builder

COPY . /auth
WORKDIR /auth

RUN go mod download
RUN go build -o ./bin/main cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /auth/bin/main .

CMD ["./main"]
