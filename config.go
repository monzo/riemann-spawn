package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/golang/glog"
	"github.com/riemann/riemann-go-client"
)

// MetricDefinition models a repeating metric
type MetricDefinition struct {
	Event         riemanngo.Event
	RatePerMinute int
}

// Configuration Format
type Configuration struct {
	RiemannHost string `json:"riemann-host"`
	RiemannPort int    `json:"riemann-port"`
	Metrics     []MetricDefinition
}

// LoadConfiguration loads the configuration from the provided filepath
func LoadConfiguration(filepath string) (Configuration, error) {
	config := Configuration{}
	file, e := ioutil.ReadFile(filepath)
	if e != nil {
		return config, e
	}
	err := json.Unmarshal(file, &config)
	return config, err
}

// GetConfig either from a provided filepath or the DefaultConfig
// filepath will either be the empty string or a value
func GetConfig(filepath string) Configuration {
	if filepath != "" {
		config, err := LoadConfiguration(filepath)
		if err != nil {
			glog.Fatalf("Couldn't Load Condifuration: %s", err)
		}
		return config
	}
	return DefaultConfig()
}
