package model

type File struct {
	DownloadUrl string `json:"downloadUrl"`
	Sha1        string `json:"sha1"`
	Name        string `json:"name"`
}

type Core struct {
	Name        string `json:"name"`
	GameVersion string `json:"gameVersion"`
	Build       string `json:"build,omitempty"`
	File        File   `json:"file"`
}

type Modification struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Website string `json:"website,omitempty"`
	File    File   `json:"file"`
}

type PackConfig struct {
	Name        string         `json:"name"`
	Version     string         `json:"version"`
	Description string         `json:"description,omitempty"`
	Author      string         `json:"author"`
	Website     string         `json:"website,omitempty"`
	Core        Core           `json:"core,omitempty"`
	Plugins     []Modification `json:"plugins,omitempty"`
	Mods        []Modification `json:"mods,omitempty"`
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
