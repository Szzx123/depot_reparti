package message

type SnapshotMessage struct {
	message string
	color   Color
}

type Color int

const (
	White Color = 0
	Red   Color = 1
)

func (c Color) String() string {
	switch c {
	case Red:
		return "rouge"
	case White:
		return "blanc"
	default:
		return "unknown"
	}
}

func (sm SnapshotMessage) Get_Color() Color {
	return sm.color
}

func (sm SnapshotMessage) Get_Message() string {
	return sm.message
}
