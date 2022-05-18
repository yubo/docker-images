# demo metrics exporter

- start demo
```
$ docker run --rm -p 9090:9090 -e DEBUG=true -it otel-demo-metrics-exporter:latest
```

- get metric
```
$ curl -sS http://0.0.0.0:9090/metrics
# TYPE demo_metrics_exporter gauge
demo_metrics_exporter{n="0"} 0.510023
demo_metrics_exporter{n="1"} -0.510023
```
