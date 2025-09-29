package system

import (
	"dream-vue-admin/global"
	"dream-vue-admin/models/system_model"
	"dream-vue-admin/server/system_server"
	"dream-vue-admin/util/constants"
	"dream-vue-admin/util/http/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DictApi 字典接口
type DictApi struct {
	server *system_server.DictServer
}

// NewDictApi 创建字典接口
func NewDictApi() *DictApi {
	return &DictApi{
		server: system_server.NewDictServer(),
	}
}

// Create 创建字典
func (api *DictApi) Create(c *gin.Context) {
	var reqData system_model.Dict
	var err error
	if err = c.ShouldBindJSON(&reqData); err != nil {
		global.Log.Warnln("【字典-创建】解析绑定结构体失败：", err.Error())
		response.Fail(c, http.StatusBadRequest, constants.RespErrorMsgBadRequest)
		return
	}
	if err = api.server.Create(c, &reqData); err != nil {
		global.Log.Warnln("【字典-创建】数据库调用失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("创建字典失败：", err.Error()))
		return
	}
	response.Success(c, reqData)
}

// Remove 删除字典
func (api *DictApi) Remove(c *gin.Context) {
	var err error
	var ids []any

	err = c.ShouldBindJSON(&ids)
	if err != nil {
		global.Log.Warnln("【字典-删除】解析绑定结构体失败：", err.Error())
		response.Fail(c, http.StatusBadRequest, constants.RespErrorMsgBadRequest)
		return
	}

	if len(ids) == 0 {
		response.Fail(c, http.StatusBadRequest, constants.RespErrorMsgBadRequest)
		return
	}

	err = api.server.Remove(c, ids)
	if err != nil {
		global.Log.Warnln("【字典-删除】数据库调用失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("删除字典失败：", err.Error()))
		return
	}
	response.SuccessWithMsg(c, constants.RespRemoveSuccessMsg)
}

// Modify 修改字典
func (api *DictApi) Modify(c *gin.Context) {
	var err error
	var data system_model.Dict
	var id = c.Param("id")
	if id == "" {
		global.Log.Warnln("【字典-修改】参数id为空")
		response.Fail(c, http.StatusBadRequest, constants.RespErrorMsgBadRequest)
		return
	}

	err = c.ShouldBindJSON(&data)
	if err != nil {
		global.Log.Warnln("【字典-修改】解析绑定结构体失败：", err.Error())
		response.Fail(c, http.StatusBadRequest, constants.RespErrorMsgBadRequest)
		return
	}

	err = api.server.Modify(c, id, &data)
	if err != nil {
		global.Log.Warnln("【字典-修改】数据库调用失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("修改字典失败：", err.Error()))
		return
	}

	response.SuccessWithMsg(c, constants.RespModifySuccessMsg)
}

// List 查询字典列表
func (api *DictApi) List(c *gin.Context) {
	var err error
	var queryData system_model.DictQuery
	err = c.ShouldBindQuery(&queryData)
	if err != nil {
		global.Log.Warnln("【字典-列表】解析绑定结构体失败：", err.Error())
	}
	list, total, err := api.server.List(&queryData)
	if err != nil {
		global.Log.Warnln("【字典-列表】数据库调用失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("获取字典列表失败：", err.Error()))
		return
	}

	response.SuccessWithList(c, list, total, &queryData.BaseQuery)
}

func (api *DictApi) GetById(c *gin.Context) {
	var err error
	var data *system_model.Dict
	var id = c.Param("id")
	if id == "" {
		global.Log.Warnln("【字典-详情】参数id为空")
		response.Fail(c, http.StatusBadRequest, constants.RespErrorMsgBadRequest)
		return
	}

	data, err = api.server.GetById(id)
	if err != nil {
		global.Log.Warnln("【字典-详情】数据库调用失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("获取字典详情失败：", err.Error()))
		return
	}

	response.Success(c, data)
}
