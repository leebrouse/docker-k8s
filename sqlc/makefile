# Makefile for running PostgreSQL container with Docker

# Docker image name
IMAGE_NAME=postgres
# Container name
CONTAINER_NAME=postgres
# PostgreSQL user and password
POSTGRES_USER=root
POSTGRES_PASSWORD=root

# Default target to start PostgreSQL container
run:
	docker run -d \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-p 5432:5432 \
		--name $(CONTAINER_NAME) $(IMAGE_NAME)

# Stop and remove the PostgreSQL container
stop:
	docker stop $(CONTAINER_NAME)

# Restart the PostgreSQL container
restart: docker restart $(CONTAINER_NAME)

# Check the status of the PostgreSQL container
status:
	docker ps -f name=$(CONTAINER_NAME)

# Remove the Docker image (if needed)
remove-image:
	docker rmi $(IMAGE_NAME)
