package initalize

import (
	"dream-vue-admin/router/v1/generate_router"
	"dream-vue-admin/router/v1/system_router"

	"github.com/gin-gonic/gin"
)

func InitV1Router(group *gin.RouterGroup) {
	//初始化生成表路由
	generate_router.InitGenerateRouter(group)
	//初始化字典路由
	system_router.InitDictRouter(group)
}
