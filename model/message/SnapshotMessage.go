package message

type SnapshotMessage struct {
	Site        string `json:"site"`         // site qui appartient l'instantané
	TypeMessage string `json:"type_message"` // "demandeSnapshot" ou "generateSnapshot"
	Horloge     string `json:"horloge"`      // horloge vectorielle du site à moment de l'instantané
	Snapshot    string `json:"snapshot"`     // information de snapshot
}

//type Color int
//
//const (
//	White Color = 0
//	Red   Color = 1
//)
//
//func (c Color) String() string {
//	switch c {
//	case Red:
//		return "rouge"
//	case White:
//		return "blanc"
//	default:
//		return "unknown"
//	}
//}

//
//func New_SnapshotMessage(site string, h int, color Color) *SnapshotMessage {
//	return &SnapshotMessage{
//		site:    site,
//		horloge: h,
//		color:   color,
//	}
//}
//
//func (sm SnapshotMessage) Get_Color() Color {
//	return sm.color
//}
//
//func (sm SnapshotMessage) Get_Site() string {
//	return sm.site
//}
