package system_model

import "dream-vue-admin/models/core"

// Dict 字典管理
type Dict struct {
	core.DreamModel
	DictName string `json:"dictName" gorm:"comment:字典名称"`
	DictCode string `json:"dictCode" gorm:"comment:字典编码"`
}

// DictQuery 字典查询参数
type DictQuery struct {
	core.BaseQuery
	DictName string `json:"dictName" form:"dictName" gorm:"comment:字典名称"`
}

// DictValue 字典值
type DictValue struct {
	core.DreamModel
	DictID uint   `json:"dictId"`
	Label  string `json:"label" gorm:"comment:标签"`
	Value  uint   `json:"dictValue" gorm:"comment:值"`
}
