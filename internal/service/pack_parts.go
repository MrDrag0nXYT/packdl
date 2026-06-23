package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"packdl/internal/config"
	"packdl/internal/model"
	"path/filepath"
)

var ErrNoLinkOrFileProviden = errors.New("No link or file providen")

func DownloadCore(client *http.Client, clientVersion string, baseDir string, core model.Core) error {
	coreName := config.Unknown
	if core.Name != "" {
		coreName = core.Name
	}

	coreGameVersion := config.Unknown
	if core.GameVersion != "" {
		coreGameVersion = core.GameVersion
	}

	coreBuild := config.Unknown
	if core.Build != "" {
		coreBuild = core.Build
	}

	fmt.Printf("Downloading '%v' version %v #%v into '%v'\n", coreName, coreGameVersion, coreBuild, baseDir)

	if core.File.DownloadUrl != "" {
		if err := runFileDownload(client, clientVersion, baseDir, core.File); err != nil {
			return err
		}
		return nil
	}

	if core.File.Name != "" && core.File.Sha1 != "" {
		targetFile := filepath.Join(baseDir, core.File.Name)

		return validateHash(targetFile, core.File.Sha1)
	}

	return ErrNoLinkOrFileProviden
}

func DownloadModifications(client *http.Client, clientVersion string, baseDir string, mods []model.Modification, modsType model.ModificationType) error {
	folderName := modsType.GetFolder()
	baseDir = filepath.Join(baseDir, folderName)

	if err := os.MkdirAll(baseDir, os.ModeAppend); err != nil {
		return fmt.Errorf("Can not create dir: %w", err)
	}

	fmt.Printf("Downloading %v into '%v'\n", folderName, baseDir)

	counter := 0

	for index, mod := range mods {
		modName := config.Unknown
		if mod.Name != "" {
			modName = mod.Name
		}

		modVersion := config.Unknown
		if mod.Version != "" {
			modVersion = mod.Version
		}

		fmt.Printf("#%v. Downloading '%v' version %v\n", index+1, modName, modVersion)

		if mod.File.DownloadUrl != "" {
			if err := runFileDownload(client, clientVersion, baseDir, mod.File); err == nil {
				counter++
				continue
			}
		}

		fmt.Println("No link providen. Download skipped")

		targetFile := filepath.Join(baseDir, mod.File.Name)
		if err := validateHash(targetFile, mod.File.Sha1); err != nil {
			fmt.Println(err)
			continue
		}

		counter++
	}

	fmt.Printf("Success: %v, %v of %v\n", folderName, counter, len(mods))

	return nil
}
