package cmd

import (
	"fmt"
	"net/http"
	"packdl/internal/config"
	"packdl/internal/model"
	"packdl/internal/service"
)

func DownloadPack(configPath string, clientVersion string) {
	packConfig, baseDir := config.LoadPackConfig(configPath)

	client := http.Client{}

	if err := service.DownloadCore(&client, clientVersion, baseDir, packConfig.Core); err != nil {
		fmt.Println(err)
	}

	if len(packConfig.Plugins) > 0 {
		if err := service.DownloadModifications(&client, clientVersion, baseDir, packConfig.Plugins, model.Plugin); err != nil {
			fmt.Println(err)
		}
	}

	if len(packConfig.Mods) > 0 {
		if err := service.DownloadModifications(&client, clientVersion, baseDir, packConfig.Mods, model.Mod); err != nil {
			fmt.Println(err)
		}
	}
}
