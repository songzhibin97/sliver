package router

import (
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

func (u *UserRouter) Register(g *gin.Engine) {
	u.engine = g
}
