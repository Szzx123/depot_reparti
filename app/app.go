package main

import (
	"flag"

	"github.com/Szzx123/depot_reparti/model/site"
	"github.com/Szzx123/depot_reparti/router"
)

func main() {
	var port = flag.String("port", "8080", "numero du port")
	var addr = flag.String("addr", "localhost", "url")
	flag.Parse()
	r := router.Router()
	site := site.New_Site(*port)
	go site.Run()
	go r.Run(*addr + ":" + *port)
	select {}
}
