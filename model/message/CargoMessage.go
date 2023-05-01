package message

type CargoMessage struct {
	Site        string `json:"site"`
	TypeMessage string `json:"type_message"`
	Stock_A     int    `json:"stock_A"`
	Stock_B     int    `json:"stock_B"`
	Stock_C     int    `json:"stock_C"`
}
