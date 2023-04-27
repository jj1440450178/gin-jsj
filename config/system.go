package config

type System struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`                // 环境值
	Addr    int    `mapstructure:"addr" json:"addr" yaml:"addr"`             // 端口值
	DbType  string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`    // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
}

var SystemCfg = System{}
