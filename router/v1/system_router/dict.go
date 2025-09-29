package system_router

import (
	"dream-vue-admin/api/v1/system"

	"github.com/gin-gonic/gin"
)

func InitDictRouter(v1group *gin.RouterGroup) {
	group := v1group.Group("/system/dict")
	api := system.NewDictApi()
	{
		group.POST("", api.Create)
		group.DELETE("", api.Remove)
		group.PUT("/:id", api.Modify)
		group.GET("/list", api.List)
		group.GET("/:id", api.GetById)
	}
}
