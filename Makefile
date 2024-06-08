# Set the default goal
.DEFAULT_GOAL := all

# Declare phony targets
.PHONY: all docker-compose build run

# Target to run Docker Compose
docker-compose:
	docker-compose up -d

# Target to compile the Go application
build:
	go build -o finapp .

# Target to start the Go application, assuming the environment variable is provided from outside
run:
	./finapp

test:
	ginkgo ./...

# Default target to run all steps
all: docker-compose build run
