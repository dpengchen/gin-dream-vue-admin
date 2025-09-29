package generate

import (
	"bytes"
	"dream-vue-admin/models/generate_model"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

//go:embed template/*
var files embed.FS

// GeneratorTemplate  生成模板
func GeneratorTemplate(table *generate_model.GenerateTable) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Println("获取当前目录错误！", err.Error())
	}

	apiPath := fmt.Sprintf("%s/api/%s/%s/%s.go", dir, table.GenerateVersion, table.GenerateBasePath, table.FolderName)
	modelPath := fmt.Sprintf("%s/models/%s/%s_model/%s.go", dir, table.GenerateVersion, table.GenerateBasePath, table.FolderName)
	serverPath := fmt.Sprintf("%s/server/%s/%s_server/%s.go", dir, table.GenerateVersion, table.GenerateBasePath, table.FolderName)
	routerPath := fmt.Sprintf("%s/router/%s/%s_router/%s.go", dir, table.GenerateVersion, table.GenerateBasePath, table.FolderName)

	if err = saveFile(table, apiPath, "back_api.template"); err != nil {
		return err
	}
	err = saveFile(table, modelPath, "back_model.template")
	if err != nil {
		return err
	}
	err = saveFile(table, routerPath, "back_router.template")
	if err != nil {
		return err
	}
	err = saveFile(table, serverPath, "back_server.template")
	if err != nil {
		return err
	}

	apiWebPath := fmt.Sprintf("%s/back_end_ui/src/api/%s/%s.ts", dir, table.GenerateBasePath, table.FileName)
	interfaceWebPath := fmt.Sprintf("%s/back_end_ui/src/interface/%s/%s.ts", dir, table.GenerateBasePath, table.FileName)
	viewWebPath := fmt.Sprintf("%s/back_end_ui/src/views/%s/%sView.vue", dir, table.GenerateBasePath, table.StructName)
	err = saveFile(table, apiWebPath, "web_api.template")
	if err != nil {
		return err
	}
	err = saveFile(table, interfaceWebPath, "web_interface.template")
	if err != nil {
		return err
	}
	err = saveFile(table, viewWebPath, "web_view.template")
	if err != nil {
		return err
	}

	return nil
}

// saveFile 保存文件
func saveFile(table *generate_model.GenerateTable, filePath, fileName string) error {

	//判断文件夹是否存在，不存在则创建
	index := strings.LastIndex(filePath, "/")
	folderPath := filePath[:index]
	_, err := os.Stat(folderPath)
	if err != nil {
		err = os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	//创建文件
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	//读取内嵌模板文件
	allData, err := files.ReadFile(fmt.Sprint("template/", fileName))
	if err != nil {
		return err
	}

	//解析文件
	tmp, err := template.New(fileName).Parse(string(allData))
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(make([]byte, 0))

	//渲染文件输出
	err = tmp.Execute(buffer, table)
	content := buffer.String()
	content = strings.ReplaceAll(content, "webTemplate-start", "{{")
	content = strings.ReplaceAll(content, "webTemplate-end", "}}")

	_, err = file.Write([]byte(content))
	return err
}
