.PHONY: build
build:
	docker build -t otelcol-geoip:latest -f ./Dockerfile .

.PHONY: push
push:
	docker tag otelcol-geoip:latest ybbbbasdf/otelcol-geoip:latest
	docker push ybbbbasdf/otelcol-geoip:latest
