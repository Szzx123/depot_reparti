package depot

type Depot struct {
	storeHouse map[string]int //key:name of goods val:number
}

func New_Depot() *Depot {
	store_house := make(map[string]int)
	return &Depot{
		storeHouse: store_house,
	}
}
