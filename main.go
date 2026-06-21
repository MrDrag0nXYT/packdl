package main

import (
	"flag"
	"fmt"
	"os"
	"packdl/internal/config"
	"packdl/internal/service"
	"packdl/internal/util"
)

var Version = "unknown_devbuild"

func main() {
	flag.Usage = showUsage
	flag.Parse()
	launchArgs := flag.Args()

	if len(launchArgs) == 0 {
		service.DownloadPack(config.DefaultConfigFileName, Version)
		return
	}

	command := launchArgs[0]
	switch command {
	case "ver", "version":
		fmt.Printf("packdl, version %v\n", Version)
		return

	case "generate", "gen":
		service.GeneratePackConfig(launchArgs)

	default:
		service.DownloadPack(command, Version)
	}

	util.ClickToExit()
}

var showUsage = func() {
	fmt.Fprintln(os.Stderr, "Usage of packdl:")
	fmt.Fprintln(os.Stderr, "  </path/to/pack> - download pack from config")
	fmt.Fprintln(os.Stderr, "  gen(generate) </path/to/pack> - generate config for pack")
	fmt.Fprintln(os.Stderr, "  ver(version) - show version")
	flag.PrintDefaults()
}
