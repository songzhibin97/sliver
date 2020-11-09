// +build windows

package server

import (
	"github.com/SliverHorn/sliver/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Init(address string, router *gin.Engine) interfaces.BaseServer {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}