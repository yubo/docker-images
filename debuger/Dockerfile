FROM debian:stretch

RUN apt-get update && apt-get install -y \
    curl \
    iputils-ping \
    trace-cmd \
    iproute2 \
    net-tools \
    mtr-tiny \
    telnet \
 && rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["bash"]
