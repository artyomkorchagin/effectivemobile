FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/subscription-service ./cmd/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/subscription-service /app/subscription-service
COPY  ./migrations /app/migrations
COPY ./docs /app/docs
COPY ./.env /app/.env
EXPOSE 3000
CMD ["/app/subscription-service"]