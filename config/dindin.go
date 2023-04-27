package config

type Dindin struct {
	AccessToken string `mapstructure:"access_token" json:"access_token" yaml:"access_token"`
	Secret      string `mapstructure:"secret" json:"secret" yaml:"secret"`
}

var DindinCfg = Dindin{}
