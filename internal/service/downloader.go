package service

import (
	"fmt"
	"packdl/internal/config"
)

func DownloadPack(configPath string) {
	packConfig := config.LoadPackConfig(configPath)

	fmt.Println(packConfig)
}
