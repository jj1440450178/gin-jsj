package app

import (
	"gin-jj/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	C     *config.Config
	Log   *zap.Logger
	Db    *gorm.DB
	Redis *redis.Client
)

type server interface {
	ListenAndServe() error
}

func InitServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
