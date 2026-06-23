package cmd

import (
	"errors"
	"fmt"
	"os"
	"packdl/internal/config"
	"packdl/internal/model"
	"packdl/internal/service/parser"
	"packdl/internal/util"
	"path/filepath"
	"strings"
)

func GeneratePackConfig(launchArgs []string, coreFile string) {
	baseDir, configFile := checkIfConfigDirExist(launchArgs)

	fmt.Printf("Generating packdl config from '%v'\n", baseDir)
	fmt.Printf("Using config file '%v'\n", configFile)

	packConfig := loadPackIfExist(configFile)

	core := parser.ParseCore(baseDir, coreFile)
	packConfig.Core = core

	plugins := parser.ParseBukkitPlugins(baseDir)
	packConfig.Plugins = plugins

	if len(plugins) == 0 && packConfig.Core == (model.Core{}) {
		fmt.Printf("No pack found in '%v'. Exiting...\n", baseDir)
		util.ClickToExit()
	}

	fmt.Printf("Saving packdl config into '%v'\n", configFile)
	config.SavePackConfig(packConfig, configFile)
}

func checkIfConfigDirExist(launchArgs []string) (baseDir string, configFile string) {
	baseDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(launchArgs) < 1 {
		return baseDir, filepath.Join(baseDir, config.DefaultConfigFileName)
	}

	targetPath := filepath.Clean(launchArgs[0])
	targetPath = strings.Trim(targetPath, `"`)

	stat, err := os.Stat(targetPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Path '%v' not exist!\n", targetPath)
		} else {
			fmt.Printf("Error on opening path '%v': %v\n", targetPath, err)
		}
		os.Exit(1)
	}

	if !stat.IsDir() {
		baseDir = filepath.Dir(targetPath)
		return baseDir, targetPath
	}

	baseDir = targetPath
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
