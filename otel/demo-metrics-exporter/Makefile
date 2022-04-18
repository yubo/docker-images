.PHONY: build
build:
	docker build -t otel-demo-metrics-exporter:latest -f ./Dockerfile .

.PHONY: push
push:
	docker tag otel-demo-metrics-exporter:latest ybbbbasdf/otel-demo-metrics-exporter:latest
	docker push ybbbbasdf/otel-demo-metrics-exporter:latest

.PHONY: run
run:
	docker run --rm -p 9090:9090 -e DEBUG=true -it otel-demo-metrics-exporter:latest
