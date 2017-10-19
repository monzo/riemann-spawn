package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"net/url"

	"github.com/golang/glog"
	"github.com/riemann/riemann-go-client"
)

// MetricDefinition models a repeating metric
type MetricDefinition struct {
	Event         riemanngo.Event
	RatePerMinute float64
}

func (m MetricDefinition) RepeatDuration() time.Duration {
	return time.Duration(float64(time.Minute) / m.RatePerMinute)
}

// Configuration Format
type Configuration struct {
	RiemannURI string
	Metrics    []MetricDefinition
}

func (c Configuration) GenerateClient() riemanngo.Client {
	url, err := url.Parse(c.RiemannURI)
	if err != nil {
		glog.Fatalf("Couldn't parse RiemannURI: %s", c.RiemannURI)
	}
	switch url.Scheme {
	case "udp":
		return riemanngo.NewUdpClient(fmt.Sprintf("%s:%s", url.Hostname(), url.Port()))
	case "tcp":
		return riemanngo.NewTcpClient(fmt.Sprintf("%s:%s", url.Hostname(), url.Port()))
	default:
		glog.Fatalf("RiemannURI must be either TCP or UDP schemed: %s", c.RiemannURI)
		return nil
	}
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
