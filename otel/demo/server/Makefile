.PHONY: build
build:
	docker build -t otel-demo-server:latest -f ./Dockerfile .

.PHONY: push
push:
	docker tag otel-demo-server:latest ybbbbasdf/otel-demo-server:latest
	docker push ybbbbasdf/otel-demo-server:latest

.PHONY: run
run:
	docker run --rm --net=host --hostname otel-demo-server -i -t otel-demo-server:latest sh

