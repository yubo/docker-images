FROM golang:1.17 as build
COPY . /usr/src/statsd/
WORKDIR /usr/src/statsd/
RUN CGO_ENABLED=0 go install ./main.go

FROM scratch
COPY --from=build /go/bin/main /main
ENV STATSD_ENDPOINT=0.0.0.0:8125
CMD ["/main"]
