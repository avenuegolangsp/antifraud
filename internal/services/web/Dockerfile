# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o webapi ./cmd/main.go


# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/webapi ./webapi
EXPOSE 8080
CMD ["./webapi"]
