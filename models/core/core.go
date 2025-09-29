package core

import "time"

type DreamModel struct {
	ID               uint       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CreateTime       *time.Time `json:"createTime" gorm:"column:create_time"`
	UpdateTime       *time.Time `json:"updateTime" gorm:"column:update_time"`
	CreateBy         uint       `json:"createBy" gorm:"column:create_by"` //创建人
	UpdateBy         uint       `json:"updateBy" gorm:"column:update_by"` //创建人
	DeleteBy         uint       `json:"deleteBy" gorm:"column:delete_by"` //删除人
	IsDel            bool       `json:"isDel" gorm:"index;column:is_del"` //软删除
	DeleteTime       *time.Time `json:"deleteTime" gorm:"column:delete_time"`
	CreateByNickname string     `json:"createByNickname" gorm:"-"` //创建人
	UpdateByNickname string     `json:"updateByNickname" gorm:"-"` //创建人
	DeleteByNickname string     `json:"deleteByNickname" gorm:"-"` //删除人
}
