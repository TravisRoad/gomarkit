package config

type Sqlite struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}
