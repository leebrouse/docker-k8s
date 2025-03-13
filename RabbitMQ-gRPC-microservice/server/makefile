# Makefile for microservices

# Variables
SERVICES = payment shopping order
PROTO_PATH = proto
GRPC_OUT = proto
RABBITMQ_URL = amqp://guest:guest@localhost:5672/

# Default target
all: start

# Generate gRPC code from protobuf
generate:
	@echo "Generating gRPC code..."
	@protoc --go_out=$(GRPC_OUT) --go-grpc_out=$(GRPC_OUT) $(PROTO_PATH)/*.proto
	@echo "gRPC code generated."

# Build all services
build: generate
	@echo "Building services..."
	@for service in $(SERVICES); do \
		echo "Building $$service..."; \
		go build -o bin/$$service ./$$service; \
	done
	@echo "All services built."

# Start RabbitMQ (ensure it's running)
start-rabbitmq:
	@echo "Starting RabbitMQ..."
	# @docker run -d --name rabbitMQ -p 5672:5672 -p 15672:15672 rabbitmq:3-management
	@echo "RabbitMQ started."

# Start all services
start: build start-rabbitmq
	@echo "Starting services..."
	@for service in $(SERVICES); do \
		echo "Starting $$service..."; \
		./bin/$$service & \
	done
	@echo "All services started."

# Stop all services
stop:
	@echo "Stopping services..."
	@for service in $(SERVICES); do \
		echo "Stopping $$service..."; \
		pkill -f "./bin/$$service"; \
	done
	@echo "All services stopped."

# Clean up
clean:
	@echo "Cleaning up..."
	@rm -rf bin/*
	@docker stop rabbitmq || true
	@docker rm rabbitmq || true
	@echo "Cleanup complete."

# Help
help:
	@echo "Available targets:"
	@echo "  all       - Build and start all services (default)"
	@echo "  generate  - Generate gRPC code from protobuf"
	@echo "  build     - Build all services"
	@echo "  start     - Start all services and RabbitMQ"
	@echo "  stop      - Stop all services"
	@echo "  clean     - Clean up binaries and stop RabbitMQ"
	@echo "  help      - Show this help message"

.PHONY: all generate build start-rabbitmq start stop clean help