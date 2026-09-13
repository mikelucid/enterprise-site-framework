.PHONY: help setup dev test lint build docker-build docker-up docker-down migrate seed tinker proto

help:
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:' Makefile | awk '{print "  make " $$1}'

setup:
	go mod download
	go mod tidy

dev:
	go run ./cmd/api-gateway serve

test:
	go test -v -cover ./...

lint:
	go fmt ./...

build:
	go build -o bin/framework ./cmd/api-gateway

docker-build:
	docker build -t yourorg/enterprise-framework:latest .

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

migrate:
	go run ./cmd/api-gateway migrate

seed:
	go run ./cmd/api-gateway seed

tinker:
	go run ./cmd/api-gateway tinker

proto:
	protoc --go_out=. --go-grpc_out=. ./api/proto/**/*.proto
