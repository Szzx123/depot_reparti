package message

type MutexMessage struct {
	site        string
	horloge     int
	typeMessage TypeMessage
	Cargo       string
	Quantity    int
	Operation   string
	Stock_A     int
	Stock_B     int
	Stock_C     int
}

type TypeMessage int

const (
	Release   TypeMessage = 0
	Request   TypeMessage = 1
	ACK       TypeMessage = 2
	demandeSC TypeMessage = 3
	startSC   TypeMessage = 4
	endSC     TypeMessage = 5
)

func (tm TypeMessage) String() string {
	switch tm {
	case Release:
		return "release"
	case Request:
		return "request"
	case ACK:
		return "ack"
	case demandeSC:
		return "demandeSC"
	case startSC:
		return "débutSC"
	case endSC:
		return "finSC"
	default:
		return "unknown"
	}
}

func New_MutexMessage(site string, h int, typ TypeMessage, cargo string, quantity int, operation string, stock_A int, stock_B int, stock_C int) *MutexMessage {
	return &MutexMessage{
		site:        site,
		horloge:     h,
		typeMessage: typ,
		Cargo:       cargo,
		Quantity:    quantity,
		Operation:   operation,
		Stock_A:     stock_A,
		Stock_B:     stock_B,
		Stock_C:     stock_C,
	}
}

func (mm MutexMessage) Get_Horloge() int {
	return mm.horloge
}

func (mm MutexMessage) Get_typeMessage() string {
	return mm.typeMessage.String()
}

func (mm MutexMessage) Get_Site() string {
	return mm.site
}
