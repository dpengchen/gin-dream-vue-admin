package generate_model

import "dream-vue-admin/models/core"

// GenerateTableQuery 生成表查询参数
type GenerateTableQuery struct {
	GenerateVersion string `json:"generateVersion" form:"generateVersion"`
	TableComment    string `json:"tableComment" form:"tableComment"`
	core.BaseQuery
}
