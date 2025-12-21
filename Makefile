# Autoload .env
ifneq (,$(wildcard .env))
include .env
export $(shell sed 's/=.*//' .env)
endif

run:
	go run ./cmd/file-service

build:
	go build -o file-service ./cmd/file-service