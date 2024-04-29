FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o weather .

FROM alpine

WORKDIR /app

COPY --from=builder /app/weather .

ENTRYPOINT ["./weather"]
