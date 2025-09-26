package initalize

import (
	"dream-vue-admin/router/v1/generate_router"

	"github.com/gin-gonic/gin"
)

func InitV1Router(group *gin.RouterGroup) {
	//初始化生成表路由
	generate_router.InitGenerateRouter(group)
}
