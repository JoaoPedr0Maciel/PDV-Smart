IMAGE_NAME=pdv_smart
CONTAINER_NAME=pdv_smart_container
NETWORK=pdv-smart_app-network
PORT=8080

DOCKER_STOP=$(shell sudo docker ps -a -q)
DOCKER_COMPOSE_UP=sudo docker compose up -d
DOCKER_BUILD=sudo docker build -t $(IMAGE_NAME) .
DOCKER_RUN=sudo docker run -d --name $(CONTAINER_NAME) --network $(NETWORK) -p $(PORT):$(PORT) $(IMAGE_NAME)

.PHONY: all clean build run

all: clean build run

clean:
	@if [ -n "$(DOCKER_STOP)" ]; then \
		sudo docker stop $(DOCKER_STOP); \
		sudo docker rm $(DOCKER_STOP); \
	fi

build:
	$(DOCKER_COMPOSE_UP)
	$(DOCKER_BUILD)

run:
	$(DOCKER_RUN)
