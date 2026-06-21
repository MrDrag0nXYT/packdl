package parser

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"packdl/internal/model"
	"packdl/internal/util"
	"path/filepath"

	"go.yaml.in/yaml/v4"
)

const pluginDataFileName = "plugin.yml"

var ErrPluginDataNotFound = errors.New(pluginDataFileName + " not found!")

func ParseBukkitPlugins(baseDir string) []model.Modification {
	baseDir = filepath.Join(baseDir, model.Plugin.GetFolder())
	fmt.Println("Reading pack plugins")

	if _, err := os.Stat(baseDir); err != nil {
		fmt.Printf("Directory 'plugins' not found, skipping\n")
		return []model.Modification{}
	}

	jars, err := listJarFilesPath(baseDir)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var plugins []model.Modification
	for _, jar := range jars {
		plugin, err := parsePluginData(jar)
		if err != nil {
			fmt.Println(err)
			continue
		}

		plugins = append(plugins, plugin)
	}

	fmt.Printf("Readed %v plugins\n", len(plugins))
	return plugins
}

func listJarFilesPath(baseDir string) ([]string, error) {
	root := os.DirFS(baseDir)

	files, err := fs.Glob(root, "*.jar")
	var jars []string
	if err != nil {
		return nil, err
	}

	for _, val := range files {
		jars = append(jars, filepath.Join(baseDir, val))
	}

	return jars, nil
}

func parsePluginData(jarPath string) (model.Modification, error) {
	zipFile, err := zip.OpenReader(jarPath)
	if err != nil {
		return model.Modification{}, err
	}

	var pluginDataFile *zip.File
	for _, f := range zipFile.File {
		if f.Name == pluginDataFileName {
			pluginDataFile = f
		}
	}

	if pluginDataFile == nil {
		return model.Modification{}, ErrPluginDataNotFound
	}

	pluginDataReader, err := pluginDataFile.Open()
	if err != nil {
		return model.Modification{}, err
	}
	defer pluginDataReader.Close()

	data, err := io.ReadAll(pluginDataReader)
	if err != nil {
		return model.Modification{}, err
	}

	zipFile.Close()

	var bukkitPluginMetadata BukkitPluginMetadata
	if err = yaml.Unmarshal(data, &bukkitPluginMetadata); err != nil {
		return model.Modification{}, err
	}

	fileHash, err := util.GetFileHash(jarPath)
	_, fileName := filepath.Split(jarPath)

	return model.Modification{
		Name:    bukkitPluginMetadata.Name,
		Version: bukkitPluginMetadata.Version,
		Website: bukkitPluginMetadata.Website,
		File: model.File{
			Name: fileName,
			Sha1: fileHash,
		},
	}, nil
}
