package main

import (
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"time"
)

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

func genMetrics(p *PeriodTime, n int) []string {
	metrics := []string{}
	t := time.Now()
	offset := time.Duration(p.Period / int64(n))

	for i := 0; i < n; i++ {
		t = t.Add(offset)
		metrics = append(metrics, fmt.Sprintf("demo_statsd_client:%f|g|#n:%d", math.Cos(p.Time(t)), i))
	}

	return metrics
}

func mainLoop() error {
	t := time.NewTicker(time.Second * 10).C
	p := &PeriodTime{
		Period: time.Hour.Nanoseconds(),
		Min:    -math.Pi,
		Max:    math.Pi,
	}

	cli, err := NewStatsdClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	for {
		select {
		case <-t:
			cli.Send(genMetrics(p, 2))
		}
	}
}

func main() {
	if err := mainLoop(); err != nil {
		fmt.Fprintf(os.Stderr, "err: %s\n", err)
		os.Exit(1)
	}
}

type PeriodTime struct {
	Period int64
	Min    float64
	Max    float64
}

func (p *PeriodTime) Time(t time.Time) float64 {
	return (float64(t.UnixNano()%p.Period)/float64(p.Period))*(p.Max-p.Min) + p.Min
}
