package api

import (
	"fmt"
	"gin-jj/app/service"
	"gin-jj/pkg/response"

	"github.com/gin-gonic/gin"
)

type Index struct{}

func IndexController() *Index {
	return &Index{}
}

func (api *Index) List(c *gin.Context) {
	response.Success(c, 200, "ok", nil)
}

func (api *Index) Dindin_send(c *gin.Context) {
	str, ok := c.GetQuery("s")
	if !ok {
		response.Fail(c, 4001, "请输入消息")
		return
	}

	err := service.DindinSend(str)
	if err != nil {
		response.Fail(c, 4002, fmt.Sprintf("发送失败：%s", err.Error()))
		return
	}

	response.Success(c, 200, fmt.Sprintf("发送成功：%s", str), nil)
}
