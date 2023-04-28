package main

import (
	"flag"

	"github.com/Szzx123/depot_reparti/model/controller"
)

func main() {
	var num = flag.String("n", "C1", "numero du controller")
	flag.Parse()
	ctl := controller.New_Controller(*num)
	go ctl.Run()
	select {}
}
