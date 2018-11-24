.PHONY: build build-app build-all

APP_NAME="go.web"
APP_TAG=latest
APP_IMAGE_NAME=$(APP_NAME)
ENV_FILE=".env"
APP_PORT=8080
REDIS_PORT=6379

GODEP="dep"

get-deps:
	@echo "${NOW} GETTING DEPENDENCIES..."
	@${GODEP} ensure -v

run:
	go run cmd/main.go

run-local:
	GO_ENV=development go run cmd/main.go

build: get-deps
	GOOS=linux go build -o ${APP_NAME} cmd/main.go

build-all: build build-app

build-app:
	echo "Building app image"
	docker build -t $(APP_IMAGE_NAME) .

run-app:
	docker run -p $(APP_PORT):$(APP_PORT) --name ${APP_NAME} --link redis:redis -it --env-file $(ENV_FILE) -e GO_ENV='development' $(APP_IMAGE_NAME):$(APP_TAG)

run-redis:
	docker run -p $(REDIS_PORT):$(REDIS_PORT) -d -it --name redis redis

run: run-redis run-app

stop:
	docker stop $(shell docker ps -a -q)
	docker rm $(shell docker ps -a -q)
