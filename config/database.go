package config

import "fmt"

type Sqlite struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}

type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Dbname   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Path, m.Port, m.Dbname, m.Config)
}

type Database struct {
	Type   string `mapstructure:"type" json:"type" yaml:"type"`
	Sqlite Sqlite `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
