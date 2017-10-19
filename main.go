package main

import (
	"flag"
	"fmt"
)

var (
	configPath string // the path to the configuration file
)

func main() {
	flag.StringVar(&configPath, "config", "", "configuration path")
	flag.Parse()

	config := GetConfig(configPath)

	fmt.Printf("%v", config)
}
