DOCKER_USER = d1y0rbek
API_GATEWAY_IMAGE = $(DOCKER_USER)/api_gateway:latest
ORDER_SERVICE_IMAGE = $(DOCKER_USER)/orderservice:latest
PRODUCT_SERVICE_IMAGE = $(DOCKER_USER)/product-service:latest

.PHONY: all build push up down clean deploy

all: build

build:
	docker compose build

push:
	docker push $(API_GATEWAY_IMAGE)
	docker push $(ORDER_SERVICE_IMAGE)
	docker push $(PRODUCT_SERVICE_IMAGE)

up:
	docker compose up -d

down:
	docker compose down

clean:
	docker system prune -f

deploy: build push up
