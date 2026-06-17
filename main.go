package main

import (
	"flag"
	"fmt"
	"packdl/internal/config"
	"packdl/internal/service"
)

var Version = "unknown_devbuild"

func main() {
	flag.Parse()
	launchArgs := flag.Args()

	if len(launchArgs) == 0 {
		service.DownloadPack(config.DefaultConfigFileName)
		return
	}

	command := launchArgs[0]
	switch command {
	case "version":
		fmt.Printf("packdl, version %v\n", Version)

	default:
		service.DownloadPack(command)
	}
}
