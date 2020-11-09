package boot

import (
	_ "github.com/SliverHorn/sliver/docs"
	"github.com/SliverHorn/sliver/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	ApiGroup := Router.Group("")
	router.InitAdminRouter(ApiGroup)
	return Router
}
