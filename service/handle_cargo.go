package service

import (
	"encoding/json"
	"github.com/Szzx123/depot_reparti/utils"
	"log"
	"net/http"

	"github.com/Szzx123/depot_reparti/model/message"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var ConnCargo *websocket.Conn

func Cargo_Handler(c *gin.Context) {
	var err error
	ConnCargo, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Upgrade WebSocket Failure: %s", err)
		return
	}
	go Cargo_Receive_Handler(c)
}

func Cargo_Receive_Handler(c *gin.Context) {
	defer ConnCargo.Close()
	for {
		// Lire le message
		_, message_json, err := ConnCargo.ReadMessage()
		if err != nil {
			log.Printf("Read Message Failure: %s", err)
			break
		}

		// Stockage temporaire du contenu du message
		message_content := make([]byte, len(message_json))
		copy(message_content, message_json)

		// Analyse de JSON
		var msg message.CargoMessage
		if err := json.Unmarshal(message_content, &msg); err != nil {
			log.Printf("Analyse JSON Failure: %s", err)
			continue
		}

		if msg.TypeMessage == "operateCargo" {
			if len(msg.Cargo) == 0 || len(msg.Type) == 0 || len(msg.Quantity) == 0 {
				log.Printf("Blank Message")
				continue
			}

			utils.Msg_send(utils.Msg_format("receiver", "C"+msg.Site[1:]) + utils.Msg_format("type", "demandeSC") + utils.Msg_format("cargo", msg.Cargo) + utils.Msg_format("operation", msg.Type) + utils.Msg_format("quantity", msg.Quantity))
			// Envoyer un message
			if err := ConnCargo.WriteMessage(websocket.TextMessage, []byte("demandeSC已收到")); err != nil {
				log.Printf("Send Message Failure: %s", err)
				break
			}
		}
	}
}

func Cargo_Send_Handler(msgCargo message.CargoMessage) {
	// Marshal the Person struct to a JSON byte slice
	jsonCargo, err := json.Marshal(msgCargo)
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}

	// Print the JSON string
	log.Println(string(jsonCargo))

	if ConnCargo != nil {
		err = ConnCargo.WriteMessage(websocket.TextMessage, jsonCargo)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("connection failed")
	}
}
