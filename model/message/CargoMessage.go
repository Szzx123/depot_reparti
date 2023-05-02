package message

type CargoMessage struct {
	Site        string `json:"site"`         // numéro de site
	TypeMessage string `json:"type_message"` // "operateCargo" ou "updateCargo"
	Cargo       string `json:"cargo"`        // type d'article
	Type        string `json:"type"`         // "in" ou "out"
	Quantity    string `json:"quantity"`     // quantité d'articles à manipuler
	Stock_A     int    `json:"stock_A"`      // quantité de A dans l'entrepôt
	Stock_B     int    `json:"stock_B"`      // quantité de B dans l'entrepôt
	Stock_C     int    `json:"stock_C"`      // quantité de C dans l'entrepôt
}
