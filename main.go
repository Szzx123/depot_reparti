package main

import (
	"github.com/Szzx123/depot_reparti/model/site"
	"github.com/Szzx123/depot_reparti/router"
)

func main() {
	r1 := router.Router()
	r2 := router.Router()
	r3 := router.Router()
	go site.Site_1.Run()
	go site.Site_2.Run()
	go site.Site_3.Run()
	go r1.Run(":8080")
	go r2.Run(":8081")
	go r3.Run(":8082")
	//阻塞主线程
	select {}
}
