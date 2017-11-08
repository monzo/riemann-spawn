package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	riemanngo "github.com/riemann/riemann-go-client"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyz")

var metricTypeList = []string{"counter", "gauge", "timing"}

type Worker struct {
	id     int
	config Configuration
	client riemanngo.Client
}

func NewWorker(id int, config Configuration) (*Worker, error) {
	client := config.GenerateClient()
	err := client.Connect(5)
	if err != nil {
		return nil, err
	}

	return &Worker{
		id:     id,
		config: config,
		client: client,
	}, nil
}

func (w *Worker) run() {
	log.Printf("Start worker  %d for %d metric(s) sent to %s", w.id, len(w.config.Metrics), w.config.RiemannURI)
	for _, metric := range w.config.Metrics {
		go func(m MetricDefinition) {
			w.repeatMetricRequest(m)
		}(metric)
	}
}
func random() string {
	return string([]rune{runes[rand.Intn(len(runes))]})
}

func metricType() string {
	return metricTypeList[rand.Intn(len(metricTypeList))]

}
func (w *Worker) repeatMetricRequest(metric MetricDefinition) {
	for {
		metric.Event.Attributes = map[string]string{
			"service":      random(),
			"service_name": random(),
			"instance":     random(),
			"host":         fmt.Sprintf("%d", w.id),
			"metric_type":  metricType(),
		}
		_, err := riemanngo.SendEvent(w.client, &metric.Event)
		if err != nil {
			log.Fatalf("%v", err)
		}
		select {
		case <-time.After(metric.RepeatDuration()):

			continue
		}
	}
}
