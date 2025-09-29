package test

import (
	"bytes"
	"dream-vue-admin/models/generate_model"
	"log"
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	files, err := template.ParseFiles("./templ/web_view.template")

	if err != nil {
		log.Println(err.Error())
	}
	file, err := os.Create("web_api.txt")
	buffer := bytes.NewBuffer(make([]byte, 0))
	if err != nil {
		log.Println(err.Error())
	}
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
	err = files.Execute(buffer, table)
	if err != nil {
		log.Println(err.Error())
	}
	content := buffer.String()
	content = strings.ReplaceAll(content, "webTemplate-start", "{{")
	content = strings.ReplaceAll(content, "webTemplate-end", "}}")
	_, _ = file.Write([]byte(content))
}
