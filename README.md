
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
git clone git@github.com:thilinajayanath/iptables-exporter.git
cd iptables-exporter
docker-compose up
```

## Flags
`-i` flag is used to pass the time interval in seconds between metric collection runs. Default is 60 seconds.

`-p` flag is used to pass the listening port. Defaul it 9445.

`-r` flag is used to pass the location of the rule file. Default is `/usr/share/iptables-exporter/rules.json`.

## Metrics available
At the moment only the number of inactive rules are available. Inactive rules mean the rules that are being checked but not available in iptables.

Metric name: `inactive_rules`

## Metric Location
Metrics are available in the loopback address of the host machine as below.
```
http://127.0.0.1:9455/metrics
```
And on `http:<ip>:9455/metrics` where `ip` is the ip address of the host machine in the docker network.

## Rules that needs to be checked
Iptables rules that needs to be checked are passed to the application as a json file in the following format.
Example is shown in the `rules.json` file in the git repo.
```
{
    "rules": [
        "INPUT -s 15.15.15.150 -j DROP",
        "INPUT -s 15.15.15.151 -j DROP",
        "INPUT -s 15.15.15.152 -j DROP"
    ]
}
```
