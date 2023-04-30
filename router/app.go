package router

import (
	"github.com/Szzx123/depot_reparti/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", service.Get_Index)
	r.GET("/operation", service.Cargo_Handler)
	r.GET("/snapshot", service.Snapshot_Handler)
	return r
}
