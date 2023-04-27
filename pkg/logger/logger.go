package logger

import (
	"fmt"
	"gin-jj/app"
	"gin-jj/pkg/utils"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(app.C.Logger.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", app.C.Logger.Director)
		_ = os.Mkdir(app.C.Logger.Director, os.ModePerm)
	}

	cores := Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if app.C.Logger.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
