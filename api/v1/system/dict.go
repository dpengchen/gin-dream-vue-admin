package system

import (
	"dream-vue-admin/global"
	"dream-vue-admin/models/v1/system_model"
	"dream-vue-admin/server/v1/system_server"
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
// @Summary 创建字典
// @Description 创建一个新的字典项
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param dict body system_model.Dict true "字典信息"
// @Success 200 {object} response.Response{data=system_model.Dict}
// @Router /system/dict [post]
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
// @Summary 删除字典
// @Description 根据ID列表删除字典项
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param ids body []uint true "字典ID列表"
// @Success 200 {object} response.Response
// @Router /system/dict [delete]
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
// @Summary 修改字典
// @Description 根据ID修改字典信息
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param id path string true "字典ID"
// @Param dict body system_model.Dict true "字典信息"
// @Success 200 {object} response.Response
// @Router /system/dict/{id} [put]
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
// @Summary 查询字典列表
// @Description 根据条件查询字典列表
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param dictName query string false "字典名称"
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(10)
// @Success 200 {object} response.Response{data=response.PageResult{list=[]system_model.Dict}}
// @Router /system/dict/list [get]
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

// GetById 获取字典详情
// @Summary 获取字典详情
// @Description 根据ID获取字典详情
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param id path string true "字典ID"
// @Success 200 {object} response.Response{data=system_model.Dict}
// @Router /system/dict/{id} [get]
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
