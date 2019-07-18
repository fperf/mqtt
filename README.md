# mqtt

This is a client of [fperf](https://github.com/shafreeck/fperf) which benchmarks MQTT

## How to use

### Build fperf with mqtt support
```
go get github.com/shafreeck/fperf/bin/fperf-build
go get github.com/fperf/mqtt
fperf-build github.com/fperf/mqtt

./fperf  -server="tls://ip:addr"  -goroutine=1 -connection=200  mqtt publish   -clientid="xx" -qos=1
```

