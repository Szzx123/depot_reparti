package main

import (
	"flag"

	"github.com/Szzx123/depot_reparti/model/controller"
)

func main() {
	var port = flag.String("port", "8080", "numero du port")
	flag.Parse()
	ctl := controller.New_Controller(*port)
	go ctl.Run()
	select {}
}
