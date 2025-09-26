package system_model

import "dream-vue-admin/models/core"

// Dict 字典管理
type Dict struct {
	core.DreamModel
}

// DictValue 字典值
type DictValue struct {
	core.DreamModel
	DictID uint   `json:"dictId"`
	Label  string `json:"label" gorm:"comment:标签"`
	Value  uint   `json:"dictValue" gorm:"comment:值"`
}
