package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type R struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Custom(c *gin.Context, statusCode, code int, msg string, data any) {
	c.JSON(statusCode, R{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func SuccessWithMsg(c *gin.Context, msg string) {
	Custom(c, http.StatusOK, http.StatusOK, msg, nil)
}
func Success(c *gin.Context, data any) {
	Custom(c, http.StatusOK, http.StatusOK, "success", data)
}

func Fail(c *gin.Context, code int, msg string) {
	Custom(c, http.StatusOK, code, msg, nil)
}

// FailWithInternalError 错误状态码500
func FailWithInternalError(c *gin.Context, msg string) {
	Fail(c, http.StatusInternalServerError, msg)
}
