package service

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"packdl/internal/model"
	"path"
	"path/filepath"
)

func runFileDownload(client *http.Client, baseDir string, file model.File) error {
	resp, err := client.Get(file.DownloadUrl)
	if err != nil {
		return fmt.Errorf("Error while downloading file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Server returned bad status code: %v", resp.StatusCode)
	}

	outName := getFileName(resp, file)
	outPath := filepath.Join(baseDir, outName)

	if _, err := os.Stat(outPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Downloading '%v'\n", outName)

			saveFile(resp, outPath)
		}

	} else {
		fmt.Printf("File '%v' already exist, skipping download\n", outName)
	}

	isHashValid, err := verifyHash(outPath, file.Sha1)

	if err != nil && errors.Is(err, ErrEmptyConfigHashsum) {
		fmt.Printf("Hashsum of file '%v' could not be checked: config hash is empty!\n", outName)
		return err
	}

	if isHashValid {
		fmt.Printf("Hashsum of file '%v' is valid!\n", outName)

	} else {
		fmt.Printf("Hashsum of file '%v' invalid! Deleting...\n", outName)
		if err := os.Remove(outPath); err != nil {
			return err
		}
	}

	return nil
}

func saveFile(resp *http.Response, filePath string) error {
	tempFilePath := filePath + ".tmp"
	out, err := os.Create(tempFilePath)

	if err != nil {
		return fmt.Errorf("Error while creating file: %w", err)
	}

	if _, err = io.Copy(out, resp.Body); err != nil {
		os.Remove(tempFilePath)
		return fmt.Errorf("Error while writing file: %w", err)
	}

	out.Close()
	return os.Rename(tempFilePath, filePath)
}

func getFileName(resp *http.Response, file model.File) string {
	if file.Name != "" {
		return file.Name
	}

	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		if _, params, err := mime.ParseMediaType(cd); err == nil {
			if fileName, ok := params["filename"]; ok {
				return fileName
			}
		}
	}

	if resp.Request.URL != nil {
		fileName := path.Base(resp.Request.URL.Path)

		if fileName != "." && fileName != "/" {
			return fileName
		}
	}

	return ""
}
