FROM golang:1.17 as build
RUN go env -w GOPROXY=direct
RUN CGO_ENABLED=0 go install github.com/open-telemetry/opentelemetry-collector-contrib/examples/demo/client@v0.47.0

FROM scratch
COPY --from=build /go/bin/client /client
ENV OTEL_EXPORTER_OTLP_ENDPOINT=0.0.0.0:4317
CMD ["/client"]
