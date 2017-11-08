package main

import (
	"flag"
	"log"
	"sync"
)

var (
	configPath string // the path to the configuration file
)

func main() {
	flag.StringVar(&configPath, "config", "", "configuration path")
	flag.Parse()
	config := GetConfig(configPath)

	var wg sync.WaitGroup
	// TODO handle worker graceful shutdown ...
	wg.Add(1)
	for i := 0; i < config.Workers; i++ {
		w, err := NewWorker(i, config)
		if err != nil {
			log.Fatalf("%v\n", err) // exits the program
		}
		w.run()
	}
	wg.Wait()
}
