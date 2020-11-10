package interfaces

import "github.com/gin-gonic/gin"

type Router interface {
	Regist(name string)
}

// IEngine: 引擎接口层
type IEngine interface {
	Engine() *gin.Engine
}

//// IRouter: 实现路由接口类
//type IRouter interface {
//	Register(IEngine)
//}

type Fs func(group *gin.RouterGroup)

type IRouter interface {
	Register(...Fs)
}
