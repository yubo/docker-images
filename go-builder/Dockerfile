FROM golang:1.18

RUN apt-get update \
 && apt-get install -y \
    ruby \
    golang \
    librados-dev \
    gogoprotobuf \
 && gem install fpm \
 && GOPROXY=https://goproxy.cn,direct go install github.com/goreleaser/goreleaser@latest \
 && rm -rf /var/lib/apt/lists/* /go/pkg/*

