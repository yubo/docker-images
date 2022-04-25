FROM golang:1.17 as build
COPY . /usr/src/main/
WORKDIR /usr/src/main/
RUN CGO_ENABLED=0 go install ./main.go

FROM scratch
COPY --from=build /go/bin/main /main
EXPOSE 8080
CMD ["/main"]
