package parser

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"io"
	"os"
	"packdl/internal/model"
	"packdl/internal/util"
	"path/filepath"
	"strings"
)

const (
	coreVersionDataFileName = "version.json"
	defaultCoreName         = "server.jar"
)

var ErrVersionDataNotFound = errors.New(coreVersionDataFileName + " not found!")

func ParseCore(baseDir string) model.Core {
	corePath := filepath.Join(baseDir, defaultCoreName)

	stat, err := os.Stat(corePath)
	if err != nil {
		return model.Core{}
	}

	bukkitCoreVersionMetadata, err := getVersionFromZip(corePath)
	if err != nil {
		return model.Core{}
	}

	fileHash, err := util.GetFileHash(corePath)
	if err != nil {
		return model.Core{}
	}

	fileName := stat.Name()
	nameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	return model.Core{
		Name:        nameWithoutExt,
		GameVersion: bukkitCoreVersionMetadata.GameVersion,
		File: model.File{
			Name: fileName,
			Sha1: fileHash,
		},
	}
}

func getVersionFromZip(corePath string) (BukkitCoreVersionMetadata, error) {
	zipFile, err := zip.OpenReader(corePath)
	if err != nil {
		return BukkitCoreVersionMetadata{}, err
	}

	var versionDataFile *zip.File
	for _, f := range zipFile.File {
		if f.Name == coreVersionDataFileName {
			versionDataFile = f
		}
	}

	if versionDataFile == nil {
		return BukkitCoreVersionMetadata{}, ErrVersionDataNotFound
	}

	versionDataReader, err := versionDataFile.Open()
	if err != nil {
		return BukkitCoreVersionMetadata{}, err
	}
	defer versionDataReader.Close()

	data, err := io.ReadAll(versionDataReader)
	if err != nil {
		return BukkitCoreVersionMetadata{}, err
	}

	zipFile.Close()

	var bukkitCoreVersionMetadata BukkitCoreVersionMetadata
	if err = json.Unmarshal(data, &bukkitCoreVersionMetadata); err != nil {
		return BukkitCoreVersionMetadata{}, err
	}

	return bukkitCoreVersionMetadata, nil
}
