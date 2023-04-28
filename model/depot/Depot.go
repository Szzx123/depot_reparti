package depot

import "fmt"

type Depot struct {
	StoreHouse map[string]int //key:name of goods val:number
}

func New_Depot() *Depot {
	depot := map[string]int{
		"A": 10,
		"B": 10,
		"C": 10,
	}
	return &Depot{
		StoreHouse: depot,
	}
}

func (d *Depot) Cargo_IN(cargo string, num int) {
	d.StoreHouse[cargo] += num
	fmt.Println(d)
}

func (d *Depot) Cargo_OUT(cargo string, num int) {
	if d.StoreHouse[cargo]-num >= 0 {
		d.StoreHouse[cargo] -= num
	}
	fmt.Println(d)
}

func (d *Depot) Set_Cargo(cargo string, num int) {
	d.StoreHouse[cargo] = num
}
