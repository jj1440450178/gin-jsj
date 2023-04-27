package config

type Redis struct {
	Enable   bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database int    `mapstructure:"database" json:"database" yaml:"database"`
}

var RedisCfg = Redis{}
