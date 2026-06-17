package model

type File struct {
	DownloadUrl string `json:"downloadUrl"`
	Sha1        string `json:"sha1"`
	Name        string `json:"name"`
}

type Core struct {
	Name        string `json:"name"`
	GameVersion string `json:"gameVersion"`
	Build       string `json:"build"`
	File        File   `json:"file"`
}

type Modification struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Website string `json:"website"`
	File    File   `json:"file"`
}

type PackConfig struct {
	Name        string         `json:"name"`
	Version     string         `json:"version"`
	Description string         `json:"description"`
	Author      string         `json:"author"`
	Website     string         `json:"website"`
	Core        Core           `json:"core"`
	Plugins     []Modification `json:"plugins"`
	Mods        []Modification `json:"mods"`
}

type ModificationType int

const (
	Plugin ModificationType = iota
	Mod
)

func (m ModificationType) GetFolder() string {
	switch m {
	case Plugin:
		return "plugins"
	case Mod:
		return "mods"
	default:
		return ""
	}
}
