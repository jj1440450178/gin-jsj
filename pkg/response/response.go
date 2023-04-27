package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Res(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Data: data,
		Msg:  msg,
		Code: code,
	})
}

// Success 正确返回
func Success(c *gin.Context, code int, msg string, data interface{}) {
	Res(c, code, msg, data)
}

// Fail失败返回
func Fail(c *gin.Context, code int, msg string) {
	Res(c, code, msg, nil)
}
