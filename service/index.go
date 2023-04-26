package service

import (
	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func Get_Index(c *gin.Context) {
	c.File("./views/index.html")
}
