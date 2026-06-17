package service

import (
	"fmt"
	"net/http"
	"packdl/internal/config"
	"packdl/internal/model"
	"path/filepath"
)

func DownloadPack(configPath string) {
	packConfig := config.LoadPackConfig(configPath)
	baseDir := filepath.Dir(configPath)

	client := http.Client{}

	if err := downloadCore(&client, baseDir, packConfig.Core); err != nil {
		fmt.Println(err)
	}

	if len(packConfig.Plugins) > 0 {
		if err := downloadModifications(&client, baseDir, packConfig.Plugins, model.Plugin); err != nil {
			fmt.Println(err)
		}
	}

	if len(packConfig.Mods) > 0 {
		if err := downloadModifications(&client, baseDir, packConfig.Mods, model.Mod); err != nil {
			fmt.Println(err)
		}
	}
}
