package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"packdl/internal/model"
	"path/filepath"
)

var ErrNoLinkProviden = errors.New("No link providen")

func downloadCore(client *http.Client, baseDir string, core model.Core) error {
	if core.File.DownloadUrl == "" {
		return ErrNoLinkProviden
	}

	if err := runFileDownload(client, baseDir, core.File); err != nil {
		return err
	}

	return nil
}

func downloadModifications(client *http.Client, baseDir string, mods []model.Modification, modsType model.ModificationType) error {
	folderName := modsType.GetFolder()
	baseDir = filepath.Join(baseDir, folderName)

	if err := os.MkdirAll(baseDir, os.ModeAppend); err != nil {
		return fmt.Errorf("Can not create dir: %w", err)
	}

	fmt.Printf("Downloading %v into '%v'\n", folderName, baseDir)

	counter := 0

	for index, mod := range mods {
		if mod.File.DownloadUrl == "" {
			return ErrNoLinkProviden
		}

		if mod.Name != "" {
			fmt.Printf("#%v. Downloading '%v'\n", index+1, mod.Name)

			if err := runFileDownload(client, baseDir, mod.File); err == nil {
				counter++
			}
		}
	}

	fmt.Printf("Success: %v, %v of %v\n", folderName, counter, len(mods))

	return nil
}
