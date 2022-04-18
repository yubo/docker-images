package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
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

type statsdClient struct {
	conn io.WriteCloser
}

func NewStatsdClient() (*statsdClient, error) {
	cli := new(statsdClient)
	if err := cli.Init(); err != nil {
		return nil, err
	}
	return cli, nil
}

func (p *statsdClient) Init() error {
	statsdEndpoint, ok := os.LookupEnv("STATSD_ENDPOINT")
	if !ok {
		statsdEndpoint = "0.0.0.0:8125"
	}

	udpAddr, err := net.ResolveUDPAddr("udp", statsdEndpoint)
	if err != nil {
		return err
	}

	if p.conn, err = net.DialUDP("udp", nil, udpAddr); err != nil {
		return err
	}

	return nil
}

func (p *statsdClient) Close() {
	if p.conn != nil {
		p.conn.Close()
		p.conn = nil
	}
}

func (p *statsdClient) Send(ss []string) error {
	for _, metric := range ss {
		_, err := fmt.Fprint(p.conn, metric)
		if err != nil {
			return err
		}
	}

	return nil
}

func mainLoop() error {
	cli, err := NewStatsdClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	rand.Seed(time.Now().UnixNano())
	offset := time.Duration(rand.Int63n(3600)) * time.Second
	p := newMetricsGenerator(2, time.Hour, offset, -math.Pi, math.Pi)

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			metrics := genMetrics(p)
			cli.Send(metrics)
			if DEBUG {
				fmt.Println(metrics)
			}
		}
	}
}

func main() {
	if err := mainLoop(); err != nil {
		fmt.Fprintf(os.Stderr, "err: %s\n", err)
		os.Exit(1)
	}
}

func genMetrics(p *metricsGenerator) []string {
	metrics := []string{}

	for i, value := range p.metrics() {
		metrics = append(metrics, fmt.Sprintf("demo_statsd_client:%f|g|#n:%d", value, i))
	}

	return metrics
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
