package main

import (
	"flag"
)

var configFile string

func main() {
	parseFlags()
}

func parseFlags() {
	flag.StringVar(&configFile, "config", "configs/app-dev.yml", "Path to configuration file")

	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}
}
