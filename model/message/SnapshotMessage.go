package message

type SnapshotMessage struct {
	Site        string `json:"site"`         // site qui appartient l'instantané
	TypeMessage string `json:"type_message"` // "demandeSnapshot" ou "generateSnapshot"
	Horloge     string `json:"horloge"`      // horloge vectorielle du site à moment de l'instantané
	Snapshot    string `json:"snapshot"`     // information de snapshot
}
