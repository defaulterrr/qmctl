package qm

type YamlConfig struct {
	Hosts []map[string]VM `yaml:"hosts"`
}
