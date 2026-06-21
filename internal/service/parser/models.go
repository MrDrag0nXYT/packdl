package parser

type BukkitPluginMetadata struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Website string `yaml:"website,omitempty"`
}

type BukkitCoreVersionMetadata struct {
	GameVersion string `json:"id"`
}
