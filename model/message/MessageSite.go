package message

type SiteMessage struct {
	TypeMessage string
	Cargo       string
	Operation   string
	Quantity    int
	Stock_A     int
	Stock_B     int
	Stock_C     int
}

func New_SiteMessage(type_message string, cargo string, operation string, quantity int, stock_A int, stock_B int, stock_C int) *SiteMessage {
	return &SiteMessage{
		TypeMessage: type_message,
		Cargo:       cargo,
		Operation:   operation,
		Quantity:    quantity,
		Stock_A:     stock_A,
		Stock_B:     stock_B,
		Stock_C:     stock_C,
	}
}
