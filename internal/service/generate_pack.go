package service

import (
	"errors"
	"fmt"
	"os"
	"packdl/internal/config"
	"packdl/internal/model"
	"packdl/internal/service/parser"
	"packdl/internal/util"
	"path/filepath"
)

func GeneratePackConfig(launchArgs []string) {
	baseDir, configFile := checkIfConfigDirExist(launchArgs)
	fmt.Printf("Generating packdl config from '%v'\n", baseDir)
	fmt.Printf("Using config file '%v'\n", configFile)

	packConfig := loadPackIfExist(configFile)

	core := parser.ParseCore(baseDir)
	packConfig.Core = core

	plugins := parser.ParseBukkitPlugins(baseDir)
	packConfig.Plugins = plugins

	if len(plugins) == 0 && packConfig.Core == (model.Core{}) {
		fmt.Printf("No pack found in '%v'. Exiting...\n", baseDir)
		util.ClickToExit()
	}

	config.SavePackConfig(packConfig, configFile)
}

func checkIfConfigDirExist(launchArgs []string) (baseDir string, configFile string) {
	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(launchArgs) < 2 {
		return baseDir, filepath.Join(baseDir, config.DefaultConfigFileName)
	}

	stat, err := os.Stat(launchArgs[1])
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Path '%v' not exist!\n", launchArgs[1])
		os.Exit(1)
	}

	if !stat.IsDir() {
		baseDir = filepath.Dir(launchArgs[1])
		return baseDir, launchArgs[1]
	}

	baseDir = launchArgs[1]
	configFile = filepath.Join(baseDir, config.DefaultConfigFileName)

	return baseDir, configFile
}

func loadPackIfExist(configFile string) model.PackConfig {
	lastPackConfig, _, err := config.LoadPackConfigIfExist(configFile)

	if err == nil {
		return model.PackConfig{
			Name:        lastPackConfig.Name,
			Description: lastPackConfig.Description,
			Author:      lastPackConfig.Author,
			Version:     lastPackConfig.Version,
			Website:     lastPackConfig.Website,
		}
	}

	return model.PackConfig{}
}
