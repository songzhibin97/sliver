package server

import (
	"github.com/SliverHorn/sliver/interfaces"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func Init(address string, router *gin.Engine) interfaces.BaseServer {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}