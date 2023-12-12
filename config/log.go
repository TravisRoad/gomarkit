package config

type Log struct {
	Dir        string `mapstructure:"dir" json:"dir" yaml:"dir"`
	Rotate     bool   `mapstructure:"rotate" json:"rotate" yaml:"rotate"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
}
