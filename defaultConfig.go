package main

import riemanngo "github.com/riemann/riemann-go-client"

// DefaultConfig : Defauly Config
func DefaultConfig() Configuration {
	return Configuration{
		RiemannURI: "tcp://127.0.0.1:5555",
		Metrics: []MetricDefinition{
			MetricDefinition{
				Event: riemanngo.Event{
					Service:    "ExampleCounter",
					Host:       "ExampleHost",
					Metric:     1,
					Attributes: map[string]string{"metric_type": "Counter"},
				},
				RatePerMinute: 60},

			MetricDefinition{
				Event: riemanngo.Event{
					Service:    "ExampleGauge",
					Host:       "ExampleHost",
					Metric:     100.49,
					Attributes: map[string]string{"metric_type": "Gauge"},
				},
				RatePerMinute: 10},
		},
	}
}
