package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"net/url"

	"github.com/riemann/riemann-go-client"
)

// MetricDefinition models a repeating metric
type MetricDefinition struct {
	Event         riemanngo.Event `json:"event"`
	RatePerSecond float64         `json:"rate_per_second"`
}

func (m MetricDefinition) RepeatDuration() time.Duration {
	return time.Duration(float64(time.Second) / m.RatePerSecond)
}

// Configuration Format
type Configuration struct {
	Workers    int                `json:"workers"`
	RiemannURI string             `json:"riemann-uri"`
	Metrics    []MetricDefinition `json:"metrics"`
}

func (c Configuration) GenerateClient() riemanngo.Client {
	url, err := url.Parse(c.RiemannURI)
	if err != nil {
		log.Fatalf("Couldn't parse RiemannURI: %s", c.RiemannURI)
	}
	switch url.Scheme {
	case "udp":
		return riemanngo.NewUdpClient(fmt.Sprintf("%s:%s", url.Hostname(), url.Port()))
	case "tcp":
		return riemanngo.NewTcpClient(fmt.Sprintf("%s:%s", url.Hostname(), url.Port()))
	default:
		log.Fatalf("RiemannURI must be either tcp or udp schemed: %s", c.RiemannURI)
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
			log.Fatal("Couldn't Load Condifuration: %s", err)
		}
		log.Printf("Processed Configuration From File: %s", filepath)
		return config
	}
	log.Printf("Using Default Configuration")
	return DefaultConfig()
}
