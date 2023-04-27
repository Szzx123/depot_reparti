package depot

import "fmt"

type Depot struct {
	storeHouse map[string]int //key:name of goods val:number
}

func New_Depot() *Depot {
	depot := map[string]int{
		"A": 10,
		"B": 10,
		"C": 10,
		"D": 10,
	}
	return &Depot{
		storeHouse: depot,
	}
}

func (d *Depot) Cargo_IN(cargo string, num int) {
	d.storeHouse[cargo] += num
	fmt.Println(d)
}

func (d *Depot) Cargo_OUT(cargo string, num int) {
	if d.storeHouse[cargo]-num >= 0 {
		d.storeHouse[cargo] -= num
	}
	fmt.Println(d)
}
