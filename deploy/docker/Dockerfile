FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/framework ./cmd/api-gateway

FROM alpine:3.18
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/bin/framework /app/framework
COPY --from=builder /app/config /app/config
EXPOSE 8000
CMD ["/app/framework", "serve"]
