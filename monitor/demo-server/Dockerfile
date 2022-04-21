FROM golang:1.17 as build
RUN go env -w GOPROXY=direct
RUN CGO_ENABLED=0 go install github.com/open-telemetry/opentelemetry-collector-contrib/examples/demo/server@v0.47.0

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
ARG USER_UID=10001
USER ${USER_UID}
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build /go/bin/server /server
ENV OTEL_EXPORTER_OTLP_ENDPOINT=0.0.0.0:4317
EXPOSE 7080
CMD ["/server"]
