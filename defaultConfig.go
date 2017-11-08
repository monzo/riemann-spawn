package main

import riemanngo "github.com/riemann/riemann-go-client"

// DefaultConfig : Defauly Config
func DefaultConfig() Configuration {
	return Configuration{
		Workers:    64,
		RiemannURI: "tcp://127.0.0.1:5555",
		Metrics: []MetricDefinition{
			MetricDefinition{
				Event: riemanngo.Event{
					Service:    "ExampleCounter",
					Host:       "ExampleHost",
					Metric:     1,
					Attributes: map[string]string{"metric_type": "Counter"},
				},
				RatePerSecond: 100,
			},
		},
	}
}
