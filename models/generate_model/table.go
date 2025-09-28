package generate_model

import "dream-vue-admin/models/core"

// GenerateTable 生成表
type GenerateTable struct {
	core.DreamModel
	GenerateVersion  string            `json:"generateVersion" gorm:"生成版本"`
	GenerateBasePath string            `json:"generateBasePath" gorm:"生成基础目录默认v1都放在v1目录下"`
	StructName       string            `json:"structName" gorm:"comment:结构体名称"`
	TableComment     string            `json:"tableComment" gorm:"comment:表描述"`
	SoftDelete       bool              `json:"softDelete" gorm:"comment:软删除"`
	GenerateColumns  []GenerateColumns `json:"generateColumns" gorm:"comment:关联字段"`
	Relation         uint              `json:"relation" gorm:"comment:表关系"`
	GenerateTableID  uint              `json:"generateTableId" gorm:"comment:关联表"`
	GenerateTable    *GenerateTable    `json:"generateTable"`
}

// GenerateColumns 生成字段
type GenerateColumns struct {
	ID              uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	GenerateTableID uint   `json:"generateTableId" gorm:"comment:关联的表id"`
	StructName      string `json:"structName" gorm:"comment:结构体字段名"`
	JsonName        string `json:"jsonName" gorm:"comment:json字段名"`
	SqlName         string `json:"sqlName" gorm:"comment:sql字段名"`
	ColumnLabel     string `json:"columnLabel" gorm:"comment:字段标签"`
	ColumnType      string `json:"columnType" gorm:"comment:字段类型"`
	InputType       string `json:"inputType" gorm:"comment:输入类型"`
	SqlType         string `json:"sqlType" gorm:"comment:sql类型"`
	DictId          uint   `json:"dictId" gorm:"comment:字典id"`
	IsEdit          bool   `json:"isEdit" gorm:"comment:是否可编辑"`
	IsExport        bool   `json:"isExport" gorm:"comment:是否导出"`
	IsShow          bool   `json:"isShow" gorm:"comment:前端是否显示字段"`
	IsQuery         bool   `json:"isQuery" gorm:"comment:是否参与过滤"`
	QueryType       string `json:"queryType" gorm:"comment:过滤条件"`
	IsSort          bool   `json:"isSort" gorm:"comment:是否可排序"`
	SortType        string `json:"sortType" gorm:"comment:排序条件类型"`
	IsRequired      bool   `json:"isRequired" gorm:"comment:是否必填"`
}
