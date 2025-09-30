package generate

import (
	"dream-vue-admin/global"
	"dream-vue-admin/models/v1/generate_model"
	"dream-vue-admin/server/v1/generate_server"
	"dream-vue-admin/util/constants"
	"dream-vue-admin/util/generate"
	"dream-vue-admin/util/http/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GeneratorTableApi 生成表接口
type GeneratorTableApi struct {
	server *generate_server.GeneratorTableServer
}

// NewGeneratorTableApi 创建生成表接口
func NewGeneratorTableApi() *GeneratorTableApi {
	return &GeneratorTableApi{
		server: generate_server.NewGeneratorTableServer(),
	}
}

// Create 创建生成表
func (api *GeneratorTableApi) Create(c *gin.Context) {
	var reqData generate_model.GenerateTable
	var err error
	if err = c.ShouldBindJSON(&reqData); err != nil {
		global.Log.Warnln("【生成表-创建】解析绑定结构体失败：", err.Error())
		response.Fail(c, http.StatusBadRequest, constants.RespErrorMsgBadRequest)
		return
	}
	if err = api.server.Create(c, &reqData); err != nil {
		global.Log.Warnln("【生成表-创建】数据库调用失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("创建生成表失败：", err.Error()))
		return
	}
	response.Success(c, reqData)
}

// Remove 删除生成表
func (api *GeneratorTableApi) Remove(c *gin.Context) {
	var err error
	var id = c.Param("id")
	if id == "" {
		global.Log.Warnln("【生成表-删除】参数id为空")
		response.Fail(c, http.StatusBadRequest, constants.RespErrorMsgBadRequest)
		return
	}

	err = api.server.Remove(c, id)
	if err != nil {
		global.Log.Warnln("【生成表-删除】数据库调用失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("删除生成表失败：", err.Error()))
		return
	}
	response.SuccessWithMsg(c, constants.RespRemoveSuccessMsg)
}

// Modify 修改生成表
func (api *GeneratorTableApi) Modify(c *gin.Context) {
	var err error
	var data generate_model.GenerateTable
	var id = c.Param("id")
	if id == "" {
		global.Log.Warnln("【生成表-修改】参数id为空")
		response.Fail(c, http.StatusBadRequest, constants.RespErrorMsgBadRequest)
		return
	}

	err = c.ShouldBindJSON(&data)
	if err != nil {
		global.Log.Warnln("【生成表-修改】解析绑定结构体失败：", err.Error())
		response.Fail(c, http.StatusBadRequest, constants.RespErrorMsgBadRequest)
		return
	}

	err = api.server.Modify(c, id, &data)
	if err != nil {
		global.Log.Warnln("【生成表-修改】数据库调用失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("修改生成表失败：", err.Error()))
		return
	}

	response.SuccessWithMsg(c, constants.RespModifySuccessMsg)
}

// List 查询生成表列表
func (api *GeneratorTableApi) List(c *gin.Context) {
	var err error
	var queryData generate_model.GenerateTableQuery
	err = c.ShouldBindQuery(&queryData)
	if err != nil {
		global.Log.Warnln("【生成表-列表】解析绑定结构体失败：", err.Error())
	}
	list, total, err := api.server.List(&queryData)
	if err != nil {
		global.Log.Warnln("【生成表-列表】数据库调用失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("获取生成表列表失败：", err.Error()))
		return
	}

	response.SuccessWithList(c, list, total, &queryData.BaseQuery)
}

func (api *GeneratorTableApi) GetById(c *gin.Context) {
	var err error
	var data *generate_model.GenerateTable
	var id = c.Param("id")
	if id == "" {
		global.Log.Warnln("【生成表-详情】参数id为空")
		response.Fail(c, http.StatusBadRequest, constants.RespErrorMsgBadRequest)
		return
	}

	data, err = api.server.GetById(id)
	if err != nil {
		global.Log.Warnln("【生成表-详情】数据库调用失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("获取生成表详情失败：", err.Error()))
		return
	}

	response.Success(c, data)
}

func (api *GeneratorTableApi) Generator(c *gin.Context) {
	id := c.Param("id")
	table, err := api.server.GetById(id)
	if err != nil {
		global.Log.Warnln("【生成表-生成】数据库调用失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("获取生成表详情失败：", err.Error()))
		return
	}
	err = generate.GeneratorTemplate(table)
	if err != nil {
		global.Log.Error("【生成表-生成】生成模板失败：", err.Error())
		response.FailWithInternalError(c, fmt.Sprint("生成模板失败：", err.Error()))
		return
	}

	response.Success(c, "生成成功")
}
