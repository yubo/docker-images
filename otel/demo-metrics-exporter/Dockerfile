FROM golang:1.17 as build
COPY . /usr/src/exporter/
WORKDIR /usr/src/exporter/
RUN CGO_ENABLED=0 go install ./main.go

FROM scratch
COPY --from=build /go/bin/main /main
ENV DEMO_METRICS_EXPORTER_ENDPOINT=0.0.0.0:9090
CMD ["/main"]
