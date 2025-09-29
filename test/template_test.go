package test

import (
	"dream-vue-admin/models/generate_model"
	"dream-vue-admin/util/generate"
	"log"
	"testing"
)

func TestTemplate(t *testing.T) {
	table := generate_model.GenerateTable{
		GenerateVersion:  "v1",
		GenerateBasePath: "system",
		StructName:       "SysUser",
		FileName:         "sysUser",
		FolderName:       "sys_user",
		TableComment:     "系统用户",
		SoftDelete:       true,
		PrivateData:      true,
		GenerateColumns: []generate_model.GenerateColumns{
			{
				ID:          1,
				StructName:  "Username",
				JsonName:    "username",
				SqlName:     "username",
				ColumnLabel: "用户名",
				ColumnType:  "string",
				TsType:      "string",
				InputType:   "text",
				SqlType:     "varchar(255)",
				DictId:      1,
				IsEdit:      true,
				IsExport:    true,
				IsShow:      true,
				IsQuery:     true,
				QueryType:   "like",
				IsSort:      true,
				SortType:    "DESC",
				IsRequired:  true,
			},
			{
				ID:          2,
				StructName:  "Password",
				JsonName:    "password",
				SqlName:     "password",
				ColumnLabel: "密码",
				ColumnType:  "string",
				TsType:      "string",
				InputType:   "password",
				SqlType:     "varchar(255)",
				DictId:      1,
				IsEdit:      true,
				IsExport:    true,
				IsShow:      true,
				IsQuery:     false,
				QueryType:   "=",
				IsSort:      false,
				SortType:    "DESC",
				IsRequired:  false,
			},
			{
				ID:          3,
				StructName:  "Nickname",
				JsonName:    "nickname",
				SqlName:     "nickname",
				ColumnLabel: "昵称",
				ColumnType:  "string",
				TsType:      "string",
				InputType:   "text",
				SqlType:     "varchar(255)",
				DictId:      1,
				IsEdit:      true,
				IsExport:    true,
				IsShow:      true,
				IsQuery:     true,
				QueryType:   "like",
				IsSort:      true,
				SortType:    "DESC",
				IsRequired:  false,
			},
		},
	}
	err := generate.GeneratorTemplate(&table)
	if err != nil {
		log.Println(err.Error())
	}
}
