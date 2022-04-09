.PHONY: build
build:
	docker build -t otel-demo-client:latest -f ./Dockerfile .

.PHONY: push
push:
	docker tag otel-demo-client:latest ybbbbasdf/otel-demo-client:latest
	docker push ybbbbasdf/otel-demo-client:latest

