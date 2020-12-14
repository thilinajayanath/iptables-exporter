FROM ubuntu:20.04

RUN set -eux;\
        mkdir /usr/share/iptables-exporter; \
        apt-get update && apt-get install -y iptables; \
        rm -rf /var/lib/apt/lists/*

COPY rules.json /usr/share/iptables-exporter/

COPY iptables-exporter /usr/local/bin

ENTRYPOINT [ "/usr/local/bin/iptables-exporter" ]