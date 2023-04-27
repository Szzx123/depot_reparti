package message

type MutexMessage struct {
	site        string
	horloge     int
	typeMessage TypeMessage
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
		return "libération"
	case Request:
		return "requête"
	case ACK:
		return "accusé"
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

func New_MutexMessage(site string, h int, typ TypeMessage) *MutexMessage {
	return &MutexMessage{
		site:        site,
		horloge:     h,
		typeMessage: typ,
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
