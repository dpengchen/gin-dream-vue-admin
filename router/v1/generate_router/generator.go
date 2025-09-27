package generate_router

import (
	"dream-vue-admin/api/v1/generate"

	"github.com/gin-gonic/gin"
)

func InitGenerateRouter(v1group *gin.RouterGroup) {
	group := v1group.Group("/generate/table")
	api := generate.NewGeneratorTableApi()
	{
		group.POST("", api.Create)
		group.DELETE("/:id", api.Remove)
		group.PUT("/:id", api.Modify)
		group.GET("/list", api.List)
		group.GET("/:id", api.GetById)
	}
}
