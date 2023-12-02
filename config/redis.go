package config

type Redis struct {
	Enable   bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
