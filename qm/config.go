package qm

type Config struct {
	// Subnet string `yaml:"subnet"`
	// Gateway string `yaml:"gateway"`
	Hosts []map[string]VM `yaml:"hosts"`
}
