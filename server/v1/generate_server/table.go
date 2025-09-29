package generate_server

import (
	"dream-vue-admin/global"
	"dream-vue-admin/models/generate_model"
	"dream-vue-admin/util/constants"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GeneratorTableServer struct {
}

func NewGeneratorTableServer() *GeneratorTableServer {
	return &GeneratorTableServer{}
}

// Create 创建生成表
func (s *GeneratorTableServer) Create(c *gin.Context, reqData *generate_model.GenerateTable) error {
	//设置创建时间
	var now = time.Now()
	reqData.CreateTime = &now
	reqData.CreateBy = c.GetUint(constants.GinContextLoginUserIdKey)
	return global.Db.Preload("GenerateColumns").Create(reqData).Error
}

// Remove 删除生成表
func (s *GeneratorTableServer) Remove(c *gin.Context, id string) error {
	var err error
	var data generate_model.GenerateTable
	err = global.Db.First(&data, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("数据不存在，不可删除！")
	}
	if err != nil {
		return err
	}

	return global.Db.Delete(&data).Error
}

// Modify 修改生成表
func (s *GeneratorTableServer) Modify(c *gin.Context, id string, data *generate_model.GenerateTable) error {
	var err error
	var count int64
	data.UpdateBy = c.GetUint(constants.GinContextLoginUserIdKey)
	now := time.Now()
	data.UpdateTime = &now

	err = global.Db.Model(&generate_model.GenerateTable{}).Where("id = ?", id).Count(&count).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || count <= 0 {
		return errors.New("数据不存在，不可修改！")
	}
	if err != nil {
		return err
	}
	db := global.Db.Begin().Debug()

	//更新生成表数据
	err = db.Model(&generate_model.GenerateTable{}).Where("id = ?", id).Select("*").Omit("create_by", "create_time").Updates(data).Error
	if err != nil {
		db.Rollback()
		return err
	}

	err = db.Where("generate_table_id = ?", id).Delete(&generate_model.GenerateColumns{}).Error
	if err != nil {
		db.Rollback()
		return err
	}
	if len(data.GenerateColumns) > 0 {
		err = db.Model(&data).Association("GenerateColumns").Replace(data.GenerateColumns)
		if err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return nil
}

// List 查询列表
func (s *GeneratorTableServer) List(queryData *generate_model.GenerateTableQuery) ([]generate_model.GenerateTable, *int64, error) {
	var list []generate_model.GenerateTable
	var total int64
	var err error
	db := global.Db.Model(&generate_model.GenerateTable{})

	if queryData.GenerateVersion != "" {
		db = db.Where("generate_version = ?", queryData.GenerateVersion)
	}

	if queryData.TableComment != "" {
		db = db.Where("table_comment like ?", "%"+queryData.TableComment+"%")
	}

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

// GetById 获取生成表通过ID
func (s *GeneratorTableServer) GetById(id string) (*generate_model.GenerateTable, error) {
	var err error
	var data generate_model.GenerateTable
	err = global.Db.Preload("GenerateColumns").First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}
