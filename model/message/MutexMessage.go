package message

type MutexMessage struct {
	horloge     int
	typeMessage TypeMessage
}

type TypeMessage int

const (
	Release   TypeMessage = 0
	Request   TypeMessage = 1
	ACK       TypeMessage = 3
	demandeSC TypeMessage = 4
	startSC   TypeMessage = 5
	endSC     TypeMessage = 6
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

func New_MutexMessage(h int, typ TypeMessage) *MutexMessage {
	return &MutexMessage{
		horloge:     h,
		typeMessage: typ,
	}
}
