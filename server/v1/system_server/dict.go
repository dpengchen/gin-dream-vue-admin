package system_server

import (
	"dream-vue-admin/global"
	"dream-vue-admin/models/system_model"
	"dream-vue-admin/util/constants"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DictServer struct {
}

func NewDictServer() *DictServer {
	return &DictServer{}
}

// Create 创建字典
func (s *DictServer) Create(c *gin.Context, reqData *system_model.Dict) error {
	//设置创建时间
	var now = time.Now()
	reqData.CreateTime = &now
	reqData.CreateBy = c.GetUint(constants.GinContextLoginUserIdKey)
	return global.Db.Create(reqData).Error
}

// Remove 删除字典
func (s *DictServer) Remove(c *gin.Context, ids []any) error {
	return global.Db.Where("id in ?", ids).Delete(&system_model.Dict{}).Error
}

// Modify 修改字典
func (s *DictServer) Modify(c *gin.Context, id string, data *system_model.Dict) error {
	var err error
	var count int64
	data.UpdateBy = c.GetUint(constants.GinContextLoginUserIdKey)
	now := time.Now()
	data.UpdateTime = &now

	err = global.Db.Model(&system_model.Dict{}).Where("id = ?", id).Count(&count).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || count <= 0 {
		return errors.New("数据不存在，不可修改！")
	}
	if err != nil {
		return err
	}

	return global.Db.Model(&system_model.Dict{}).Where("id = ?", id).Select("*").Omit("id", "create_by", "create_time", "delete_by", "delete_time").Updates(data).Error
}

// List 查询列表
func (s *DictServer) List(queryData *system_model.DictQuery) ([]system_model.Dict, *int64, error) {
	var list []system_model.Dict
	var total int64
	var err error
	db := global.Db.Model(&system_model.Dict{})

	//条件查询
	if queryData.DictName != "" {
		db = db.Where("dict_name like ?", fmt.Sprint('%', queryData.DictName, '%'))
	}
	//只能查询到没有删除的数据，只有在私有数据可以进行
	err = db.Count(&total).Error
	if err != nil {
		return nil, nil, err
	}

	err = db.Offset((queryData.Page - 1) * queryData.Size).Order("id DESC").Find(&list).Error
	if err != nil {
		return nil, nil, err
	}

	return list, &total, nil
}

// GetById 获取字典通过ID
func (s *DictServer) GetById(id string) (*system_model.Dict, error) {
	var err error
	var data system_model.Dict
	err = global.Db.First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}
