package routes

import (
	"gin-jj/app"
	"gin-jj/pkg/response"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	switch app.C.System.Env {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
		// 禁用 gin 输出接口访问日志
		gin.DefaultWriter = ioutil.Discard
	}

	r := gin.Default()

	//为文件提供静态地址
	r.StaticFS(app.C.Local.StorePath, http.Dir(app.C.Local.StorePath))

	//加载api路由
	ApiRoute(r)

	r.NoRoute(func(c *gin.Context) {
		response.Fail(c, 404, "没有该路由")
	})

	return r
}
