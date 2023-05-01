package message

type CargoMessage struct {
	Site        string `json:"site"` //num of site
	TypeMessage string `json:"type_message"`
	Cargo       string `json:"cargo"` // cargo name
	Type        string `json:"type"`  // in or out
	Quantity    string `json:"quantity"`
	Stock_A     int    `json:"stock_A"`
	Stock_B     int    `json:"stock_B"`
	Stock_C     int    `json:"stock_C"`
}
