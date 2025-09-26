package core

import "time"

type DreamModel struct {
	ID         uint       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	CreateBy   int        `json:"createBy"`           //创建人
	UpdateBy   int        `json:"updateBy"`           //创建人
	IsDel      int        `json:"isDel" gorm:"index"` //软删除
	DeleteBy   int        `json:"deleteBy"`           //删除人
	DeleteTime *time.Time `json:"deleteTime"`
}
