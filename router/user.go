package router

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
    engine gin.Engine
}

func (u *UserRouter) Regist(name string) {
	u.engine.Group(name)
}


