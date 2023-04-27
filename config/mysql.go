package config

import "time"

type Mysql struct {
	Enable       bool          `mapstructure:"enable" json:"enable" yaml:"enable"`
	Host         string        `mapstructure:"host" json:"host" yaml:"host"`
	Username     string        `mapstructure:"username" json:"username" yaml:"username"`
	Password     string        `mapstructure:"password" json:"password" yaml:"password"`
	Port         uint16        `mapstructure:"port" json:"port" yaml:"port"`
	Database     string        `mapstructure:"database" json:"database" yaml:"database"`
	Charset      string        `mapstructure:"charset" json:"charset" yaml:"charset"`
	Prefix       string        `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	MaxIdleConns int           `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns int           `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	MaxLifetime  time.Duration `mapstructure:"max_lifetime" json:"max_lifetime" yaml:"max_lifetime"`
}

var MsqlCfg = Mysql{}
