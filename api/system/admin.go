package system

import (
	"github.com/SliverHorn/sliver/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAdmin(c *gin.Context) {
	id := c.Param("id")
	admin := service.FindAdmin(id)
	c.JSON(http.StatusOK, admin)
}