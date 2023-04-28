package main

import (
	"flag"

	"github.com/Szzx123/depot_reparti/model/site"
	"github.com/Szzx123/depot_reparti/router"
)

func main() {
	var port = flag.String("p", "8080", "numero du port")
	var num = flag.String("n", "A1", "numero de l'application")
	flag.Parse()
	r := router.Router()
	site := site.New_Site(*num)
	go r.Run(":" + *port)
	go site.Run()
	select {}
}
