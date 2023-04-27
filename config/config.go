package config

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	System System `json:"system" yaml:"system"`
	Local  Local  `json:"local" yaml:"local"`
	Mysql  Mysql  `json:"mysql" yaml:"mysql"`
	Redis  Redis  `json:"redis" yaml:"redis"`
	Logger Logger `json:"logger" yaml:"logger"`
	Dindin Dindin `json:"dindin" yaml:"dindin"`
}

var C = &Config{
	System: SystemCfg,
	Local:  LocalCfg,
	Mysql:  MsqlCfg,
	Redis:  RedisCfg,
	Logger: LoggerCfg,
	Dindin: DindinCfg,
}

var once sync.Once

func InitConfig() *Config {
	once.Do(func() {
		// 加载 .yaml 配置
		v := viper.New()
		v.SetConfigFile("config.yaml")
		v.SetConfigType("yaml")
		err := v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		v.WatchConfig()

		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.Name)
			if err = v.Unmarshal(&C); err != nil {
				fmt.Println(err)
			}
		})
		if err = v.Unmarshal(&C); err != nil {
			fmt.Println(err)
		}
	})
	return C
}
