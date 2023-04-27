package main

import (
	"flag"
	"fmt"
	"gin-jj/app"
	"gin-jj/command"
	"gin-jj/config"
	"gin-jj/data"
	"gin-jj/pkg/logger"
	"gin-jj/routes"
)

func main() {
	var Conmand string
	flag.StringVar(&Conmand, "c", "", "输入command命令")
	flag.Parse()

	//加载config.yaml配置
	app.C = config.InitConfig()

	//加载logger日志
	app.Log = logger.InitLogger()

	//初始化mysql数据库
	data.InitMysql()

	if Conmand != "" {
		//运行cmd命令
		command.Run(Conmand)
	} else {
		//运行http服务
		r := routes.InitRoute()
		address := fmt.Sprintf(":%d", app.C.System.Addr)
		s := app.InitServer(address, r)
		fmt.Printf("运行地址:http://127.0.0.1%s\n", address)
		app.Log.Error(s.ListenAndServe().Error())
	}
}
