package generate

import (
	"bytes"
	"dream-vue-admin/models/generate_model"
	"embed"
	"errors"
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

	//后端代码生成
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

	//前端代码生成
	apiWebPath := fmt.Sprintf("%s/back_end_ui/src/api/%s/%s.ts", dir, table.GenerateBasePath, table.FolderName)
	interfaceWebPath := fmt.Sprintf("%s/back_end_ui/src/interface/%s/%s.ts", dir, table.GenerateBasePath, table.FolderName)
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

	//注册后端路由和gorm
	return registerRouterAndGorm(table, dir)

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

func registerRouterAndGorm(table *generate_model.GenerateTable, dir string) error {
	var err error

	//TODO 注册到路由数据库

	//注册到路由文件中
	err = writerFile(
		fmt.Sprintf("%s/%s", dir, "router/initalize/v1.go"),
		fmt.Sprintf("\t//初始化%s路由\n\t%s_router.Init%sRouter(group)", table.TableComment, table.GenerateBasePath, table.StructName),
		11,
	)
	if err != nil {
		return err
	}

	//注册到gorm数据库中
	return writerFile(
		fmt.Sprintf("%s/%s", dir, "pkg/gorm/gorm.go"),
		fmt.Sprintf("\t\t//注册%s模型\n\t\t&%s_model.%s{},", table.TableComment, table.GenerateBasePath, table.StructName),
		54,
	)
}

func writerFile(filePath string, writerContent string, line int) error {

	content, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New("打开路由文件错误" + filePath + err.Error())
	}
	routerFile, _ := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0666)
	defer routerFile.Close()

	//替换文件
	lines := strings.Split(string(content), "\n")
	contentLines := make([]string, len(lines)+1)
	line--
	for i := range lines {
		if i == line {
			contentLines[i] = writerContent
			contentLines[i+1] = lines[i]
			continue
		}
		if i > line {
			contentLines[i+1] = lines[i]
			continue
		}

		contentLines[i] = lines[i]
	}

	//清空routerFile，写入contentStr
	_ = routerFile.Truncate(0)
	writerContent = strings.Trim(strings.Join(contentLines, "\n"), "\x00")
	_, err = routerFile.WriteString(writerContent)
	return err
}
