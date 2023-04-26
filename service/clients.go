package service

import "github.com/Szzx123/depot_reparti/model/client"

var (
	Client_1                           = client.New_Client("8080")
	Client_2                           = client.New_Client("8081")
	Client_3                           = client.New_Client("8082")
	Clients  map[string]*client.Client = map[string]*client.Client{
		"8080": Client_1,
		"8081": Client_2,
		"8082": Client_3,
	}
)
