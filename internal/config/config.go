package config

import (
	"encoding/json"
	"fmt"
	"os"
	"packdl/internal/model"
	"packdl/internal/util"
)

func LoadPackConfig(configPath string) model.PackConfig {
	data := openFile(configPath)
	packConfig := parseFile(data)

	if !validateConfig(packConfig) {
		util.ClickToExit()
	}

	return packConfig
}

func openFile(configPath string) []byte {
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		os.Exit(1)
	}
	return data
}

func parseFile(data []byte) model.PackConfig {
	var packConfig model.PackConfig

	if err := json.Unmarshal(data, &packConfig); err != nil {
		fmt.Println("An error occured while reading file! Is config file correct?")
		os.Exit(1)
	}

	return packConfig
}

func validateConfig(packConfig model.PackConfig) bool {
	if packConfig.Name == "" || packConfig.Author == "" || packConfig.Version == "" {
		fmt.Println("Config metadata is empty!")
		return false
	}

	isCoreData := packConfig.Core.Name == "" && packConfig.Core.Build == "" && packConfig.Core.GameVersion == ""
	isModificationsData := packConfig.Plugins == nil && packConfig.Mods == nil

	if isModificationsData {
		fmt.Println("There are no plungins or mods! Have you lost something?")
	}

	if isCoreData && packConfig.Core.File == (model.File{}) {
		fmt.Println("There is no pack core metadata! Have you lost something?")
	}

	if isCoreData && isModificationsData {
		fmt.Println("There is nothing to do")
		return false
	}

	return true
}
