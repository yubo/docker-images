DOCKER_IMAGE_REPO ?= ybbbbadsf/docker-images:latest
DOCKER_RUN_ARGS ?= --rm -it

all: build

.PHONY: build
build:
	docker build -t $(DOCKER_IMAGE_REPO) -f ./Dockerfile .

.PHONY: push
push:
	docker push $(DOCKER_IMAGE_REPO)

.PHONY: run
run:
	docker run $(DOCKER_RUN_ARGS) $(DOCKER_IMAGE_REPO)
