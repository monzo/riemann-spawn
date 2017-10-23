package main

import (
	"flag"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/glog"
	riemanngo "github.com/riemann/riemann-go-client"
)

var (
	configPath string // the path to the configuration file
)

func main() {
	flag.StringVar(&configPath, "config", "", "configuration path")
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	config := GetConfig(configPath)
	glog.Info(spew.Sdump(config))

	var wg sync.WaitGroup
	for _, metric := range config.Metrics {
		wg.Add(1)
		go func(m MetricDefinition) {
			defer wg.Done()
			repeatMetricRequest(m, config.GenerateClient())
		}(metric)
	}
	wg.Wait()
}

func repeatMetricRequest(metric MetricDefinition, client riemanngo.Client) {
	for {
		go func(event riemanngo.Event) {
			glog.Info(event)

			err := client.Connect(5)
			if err != nil {
				glog.Errorf("%v", err)
				return
			}

			riemanngo.SendEvent(client, &event)
			return
		}(metric.Event)
		time.Sleep(metric.RepeatDuration())
	}
}
