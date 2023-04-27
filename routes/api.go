package routes

import (
	"gin-jj/app/controller/api"

	"github.com/gin-gonic/gin"
)

func ApiRoute(r *gin.Engine) {
	apiRoute := r.Group("/api")
	{
		index := api.IndexController()
		apiRoute.GET("/", index.List)

		apiRoute.GET("/dindin-send", index.Dindin_send)
	}
}
