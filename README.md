
# iptables-exporter

## Introduction

This is a prometheus exporter for iptables to check given set of rules exist in iptables.

## Usage
### As a a standalone application
```
git clone git@github.com:thilinajayanath/iptables-exporter.git
cd iptables-exporter
go build
sudo ./iptables-exporter -r rules.json
```
### Run it on docker

```
docker run --net=host --cap-add=NET_ADMIN --cap-add=NET_RAW --mount type=bind,source="$(pwd)"/rules.json,target=/usr/share/iptables-exporter/rules.json,readonly thilinajayanath/iptables-exporter
```
### Run it with docker-compose
```
docker-compose up
```
## Flags
`-i` flag is used to pass the time interval in seconds between metric collection runs. Default is 60 seconds.

`-p` flag is used to pass the listening port. Defaul it 9445.

`-r` flag is used to pass the location of the rule file. Default is `/usr/share/iptables-exporter/rules.json`.
