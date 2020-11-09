package router

import (
	"github.com/SliverHorn/sliver/api/system"
	"github.com/gin-gonic/gin"
)

func InitAdminRouter(Router *gin.RouterGroup)  {
	AdminRouter := Router.Group("admin")
	{
		AdminRouter.GET("/:id", system.FindAdmin)
	}
}
