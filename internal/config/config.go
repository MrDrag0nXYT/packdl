package config

import (
	"encoding/json"
	"fmt"
	"os"
	"packdl/internal/model"
	"packdl/internal/util"
	"path/filepath"
)

func LoadPackConfigIfExist(configPath string) (model.PackConfig, string, error) {
	data, baseDir, err := openFile(configPath)
	if err != nil {
		return model.PackConfig{}, "", err
	}
	packConfig := parseFile(data)

	return packConfig, baseDir, nil
}

func LoadPackConfig(configPath string) (model.PackConfig, string) {
	packConfig, baseDir, err := LoadPackConfigIfExist(configPath)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		os.Exit(1)
	}

	if !validateConfig(packConfig) {
		util.ClickToExit()
	}

	return packConfig, baseDir
}

func SavePackConfig(packConfig model.PackConfig, configPath string) {
	data := serializeConfig(packConfig)
	saveFile(data, configPath)
}

func openFile(configPath string) ([]byte, string, error) {
	stat, err := os.Stat(configPath)
	if err != nil {
		return nil, "", err
	}

	baseDir := filepath.Dir(configPath)

	if stat.IsDir() {
		baseDir = configPath
		configPath = filepath.Join(configPath, DefaultConfigFileName)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, "", err
	}

	return data, baseDir, nil
}

func parseFile(data []byte) model.PackConfig {
	var packConfig model.PackConfig

	if err := json.Unmarshal(data, &packConfig); err != nil {
		fmt.Println("An error occured while reading file! Is config file correct?")
		os.Exit(1)
	}

	return packConfig
}

func serializeConfig(packConfig model.PackConfig) []byte {
	data, err := json.Marshal(packConfig)
	if err != nil {
		fmt.Println("An error occured while parsing config for save!")
		os.Exit(1)
	}

	return data
}

func saveFile(data []byte, configPath string) error {
	return os.WriteFile(configPath, data, 0666)
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
