FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/backend ./cmd/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/backend /app/backend
EXPOSE 3000
CMD ["/app/backend"]