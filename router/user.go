package router

import (
	"github.com/SliverHorn/sliver/interfaces"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	engine *gin.Engine
}

//func (u *UserRouter) Regist(name string) {
//	u.engine.Group(name)
//}

func (u *UserRouter) Engine() *gin.Engine {
	return u.engine
}

func (u *UserRouter) Register(i interfaces.IEngine) {
	u.engine = i.Engine()
}
