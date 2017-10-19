package main

import (
	"flag"
	"sync"
	"time"

	"github.com/golang/glog"
	riemanngo "github.com/riemann/riemann-go-client"
)

var (
	configPath string // the path to the configuration file
)

func repeatMetricRequest(metric MetricDefinition) {
	for {
		go func(event riemanngo.Event) {
			glog.Info(event)
		}(metric.Event)
		<-time.After(metric.repeatDuration())
	}

}
func main() {
	flag.StringVar(&configPath, "config", "", "configuration path")
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	config := GetConfig(configPath)

	var wg sync.WaitGroup
	for _, metric := range config.Metrics {
		wg.Add(1)
		go func(m MetricDefinition) {
			defer wg.Done()
			repeatMetricRequest(m)
		}(metric)
	}
	wg.Wait()
}
