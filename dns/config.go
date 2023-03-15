package dns

type Config struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
	Timeout uint32 `yaml:"timeout"`
}
