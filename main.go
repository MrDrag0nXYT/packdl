package main

import (
	"flag"
	"fmt"
	"os"
	"packdl/internal/cmd"
	"packdl/internal/config"
	"packdl/internal/util"
)

var Version = "unknown_devbuild"

func main() {
	flag.Usage = showUsage

	if len(os.Args) == 0 {
		cmd.DownloadPack(config.DefaultConfigFileName, Version)
		return
	}

	command := os.Args[1]
	switch command {
	case "ver", "version":
		fmt.Printf("packdl, version %v\n", Version)
		return

	case "generate", "gen":
		genCmd := flag.NewFlagSet("gen", flag.ExitOnError)
		coreFile := genCmd.String("core", "", "core.jar file")

		genCmd.Parse(os.Args[2:])
		launchArgs := genCmd.Args()

		cmd.GeneratePackConfig(launchArgs, *coreFile)

	default:
		cmd.DownloadPack(command, Version)
	}

	util.ClickToExit()
}

func showUsage() {
	fmt.Fprintln(os.Stderr, "Usage of packdl:")
	fmt.Fprintln(os.Stderr, "  </path/to/pack> - download pack from config")
	fmt.Fprintln(os.Stderr, "  gen(generate) - generate config for pack")
	fmt.Fprintln(os.Stderr, "    -core core.jar - set server core jar file")
	fmt.Fprintln(os.Stderr, "    </path/to/pack>")
	fmt.Fprintln(os.Stderr, "  ver(version) - show version")
	flag.PrintDefaults()
}
