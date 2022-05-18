FROM golang:1.17 as build
COPY . /usr/src/main/
WORKDIR /usr/src/main/
RUN CGO_ENABLED=0 go install ./main.go

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
ARG USER_UID=10001
USER ${USER_UID}
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build /go/bin/main /main
ENV STATSD_ENDPOINT=0.0.0.0:8125
CMD ["/main"]
