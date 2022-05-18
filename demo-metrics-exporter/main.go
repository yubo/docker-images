package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	DEBUG = false
)

func init() {
	if os.Getenv("DEBUG") != "" {
		DEBUG = true
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	offset := time.Duration(rand.Int63n(3600)) * time.Second
	p := newMetricsGenerator(2, time.Hour, offset, -math.Pi, math.Pi)

	// http
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "Content-Type: text/plain; version=0.0.4; charset=utf-8")
		fmt.Fprintf(w, "# HELP demo_metrics_exporter demo gauge\n")
		fmt.Fprintf(w, "# TYPE demo_metrics_exporter gauge\n")
		for i, value := range p.metrics() {
			fmt.Fprintf(w, "demo_metrics_exporter{n=\"%d\"} %f \n", i, value)
			if DEBUG {
				fmt.Printf("demo_metrics_exporter{n=\"%d\"} %f \n", i, value)
			}
		}
	})

	addr, ok := os.LookupEnv("DEMO_METRICS_EXPORTER_ENDPOINT")
	if !ok {
		addr = "0.0.0.0:9090"
	}

	fmt.Printf("listening on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("exit %s\n", err)
	}

}

func newMetricsGenerator(n int, period, offset time.Duration, min, max float64) *metricsGenerator {
	return &metricsGenerator{
		n:      n,
		Period: period.Nanoseconds(),
		offset: offset,
		min:    min,
		max:    max,
		step:   period / time.Duration(n),
	}
}

type metricsGenerator struct {
	n      int
	Period int64
	offset time.Duration
	min    float64
	max    float64
	step   time.Duration
}

func (p *metricsGenerator) metrics() []float64 {
	metrics := []float64{}

	t := time.Now().Add(p.offset)
	for i := 0; i < p.n; i++ {
		t = t.Add(p.step)
		metrics = append(metrics, p.metric(t))
	}

	return metrics
}

func (p *metricsGenerator) metric(t time.Time) float64 {
	return math.Cos((float64(t.UnixNano()%p.Period)/float64(p.Period))*(p.max-p.min) + p.min)
}
