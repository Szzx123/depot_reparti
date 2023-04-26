package message

type Message struct {
	Site     string `json:"site"`  //num of site
	Cargo    string `json:"cargo"` // cargo name
	Type     string `json:"type"`  // in or out
	Quantity string `json:"quantity"`
}
