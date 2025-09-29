package response

import (
	"dream-vue-admin/models/core"
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
	Custom(c, http.StatusOK, http.StatusOK, "操作成功！", data)
}

func Fail(c *gin.Context, code int, msg string) {
	Custom(c, http.StatusOK, code, msg, nil)
}

// FailWithInternalError 错误状态码500
func FailWithInternalError(c *gin.Context, msg string) {
	Fail(c, http.StatusInternalServerError, msg)
}

// SuccessWithList 响应成功返回列表数据
func SuccessWithList(c *gin.Context, list any, total *int64, query *core.BaseQuery) {
	totalPage := *total / int64(query.Size)
	if *total%int64(query.Size) != 0 {
		totalPage++
	}
	Success(c, gin.H{
		"list":      list,
		"total":     total,
		"totalPage": totalPage,
		"page":      query.Page,
		"size":      query.Size,
	})
}
