.PHONY: build
build:
	docker build -t otel-demo-statsd:latest -f ./Dockerfile .

.PHONY: push
push:
	docker tag otel-demo-statsd:latest ybbbbasdf/otel-demo-statsd:latest
	docker push ybbbbasdf/otel-demo-statsd:latest

.PHONY: run
run:
	docker run --rm --net=host -it otel-demo-statsd:latest
