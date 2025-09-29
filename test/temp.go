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

type SysUserServer struct {
}

func NewSysUserServer() *SysUserServer {
	return &SysUserServer{}
}

// Create 创建系统用户
func (s *SysUserServer) Create(c *gin.Context, reqData *system_model.SysUser) error {
	//设置创建时间
	var now = time.Now()
	reqData.CreateTime = &now
	reqData.CreateBy = c.GetUint(constants.GinContextLoginUserIdKey)
	return global.Db.Create(reqData).Error
}

// Remove 删除系统用户
func (s *SysUserServer) Remove(c *gin.Context, ids []any) error {
	return global.Db.Where("id in ?", ids).Delete(&system_model.SysUser{}).Error
}

// Modify 修改系统用户
func (s *SysUserServer) Modify(c *gin.Context, id string, data *system_model.SysUser) error {
	var err error
	var count int64
	data.UpdateBy = c.GetUint(constants.GinContextLoginUserIdKey)
	now := time.Now()
	data.UpdateTime = &now

	userId := c.GetUint(constants.GinContextLoginUserIdKey)

	err = global.Db.Model(&system_model.SysUser{}).Where("id = ?", id).Where("create_by = ?", userId).Count(&count).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || count <= 0 {
		return errors.New("数据不存在，不可修改！")
	}
	if err != nil {
		return err
	}

	return global.Db.Model(&system_model.SysUser{}).Where("id = ?", id).Select("*").Omit("id", "create_by", "create_time", "delete_by", "delete_time").Updates(data).Error
}

// List 查询列表
func (s *SysUserServer) List(c *gin.Context, queryData *system_model.SysUserQuery) ([]system_model.SysUser, *int64, error) {
	var list []system_model.SysUser
	var total int64
	var err error
	db := global.Db.Model(&system_model.SysUser{})

	userId := c.GetUint(constants.GinContextLoginUserIdKey)

	//条件查询

	// 用户名
	if queryData.Username != null {
		db = db.Where("username = ?", queryData.Username)
	}

	// 昵称
	if queryData.Nickname != null {
		db = db.Where("nickname = ?", queryData.Nickname)
	}

	//只能查询到没有删除的数据，只有在私有数据可以进行
	err = db.Where("create_by = ?", userId).Count(&total).Error
	if err != nil {
		return nil, nil, err
	}

	err = db.Where("create_by = ?", userId).Offset((queryData.Page - 1) * queryData.Size).Order("id DESC").Find(&list).Error
	if err != nil {
		return nil, nil, err
	}

	return list, &total, nil
}

// GetById 获取系统用户通过ID
func (s *SysUserServer) GetById(c *gin.Context, id string) (*system_model.SysUser, error) {
	var err error
	var data system_model.SysUser

	userId := c.GetUint(constants.GinContextLoginUserIdKey)

	err = global.Db.Where("create_by = ?", userId).First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}
