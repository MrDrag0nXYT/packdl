package service

import (
	"fmt"
	"net/http"
	"packdl/internal/config"
	"packdl/internal/model"
)

func DownloadPack(configPath string, clientVersion string) {
	packConfig, baseDir := config.LoadPackConfig(configPath)

	client := http.Client{}

	if err := downloadCore(&client, clientVersion, baseDir, packConfig.Core); err != nil {
		fmt.Println(err)
	}

	if len(packConfig.Plugins) > 0 {
		if err := downloadModifications(&client, clientVersion, baseDir, packConfig.Plugins, model.Plugin); err != nil {
			fmt.Println(err)
		}
	}

	if len(packConfig.Mods) > 0 {
		if err := downloadModifications(&client, clientVersion, baseDir, packConfig.Mods, model.Mod); err != nil {
			fmt.Println(err)
		}
	}
}
