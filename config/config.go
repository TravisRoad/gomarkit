package config

type Config struct {
	Port   int    `mapstructure:"port" json:"port" yaml:"port"`
	Token  string `mapstructure:"token" json:"token" yaml:"token"`
	Salt   string `mapstructure:"salt" json:"salt" yaml:"salt"`
	Sqlite Sqlite `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
}
