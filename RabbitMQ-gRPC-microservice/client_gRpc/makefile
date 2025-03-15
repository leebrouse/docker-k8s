# Makefile for RabbitMQ and gRPC microservices

# Variables
SERVICES = payment shopping order
PROTO_PATH = proto
GRPC_OUT = proto
RABBITMQ_CONTAINER = rabbitmq
RABBITMQ_IMAGE = rabbitmq:3-management
RABBITMQ_PORT = 5672
RABBITMQ_MANAGEMENT_PORT = 15672

# Default target
all: start

# Generate gRPC code from protobuf
# generate:
# 	@echo "Generating gRPC code..."
# 	@protoc --go_out=$(GRPC_OUT) --go-grpc_out=$(GRPC_OUT) $(PROTO_PATH)/*.proto
# 	@echo "gRPC code generated."

# Build all services
 build: generate
	@echo "Building services..."
	@for service in $(SERVICES); do \
		echo "Building $$service..."; \
		go build -o bin/$$service ./$$service; \
	done
	@echo "All services built."

# Start RabbitMQ container
start-rabbitmq:
	@echo "Starting RabbitMQ container..."
	@if [ -z "$(shell docker ps -q -f name=$(RABBITMQ_CONTAINER))" ]; then \
		docker run -d --name $(RABBITMQ_CONTAINER) -p $(RABBITMQ_PORT):5672 -p $(RABBITMQ_MANAGEMENT_PORT):15672 $(RABBITMQ_IMAGE); \
		echo "RabbitMQ started."; \
	else \
		echo "RabbitMQ is already running."; \
	fi

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

# Stop RabbitMQ container
stop-rabbitmq:
	@echo "Stopping RabbitMQ container..."
	@docker stop $(RABBITMQ_CONTAINER) || true
	@docker rm $(RABBITMQ_CONTAINER) || true
	@echo "RabbitMQ stopped."

# Clean up
clean: stop stop-rabbitmq
	@echo "Cleaning up..."
	@rm -rf bin/*
	@echo "Cleanup complete."

# Help
help:
	@echo "Available targets:"
	@echo "  all         - Build and start all services (default)"
	@echo "  generate    - Generate gRPC code from protobuf"
	@echo "  build       - Build all services"
	@echo "  start       - Start all services and RabbitMQ"
	@echo "  stop        - Stop all services"
	@echo "  stop-rabbitmq - Stop RabbitMQ container"
	@echo "  clean       - Clean up binaries and stop RabbitMQ"
	@echo "  help        - Show this help message"

.PHONY: all generate build start-rabbitmq start stop stop-rabbitmq clean help