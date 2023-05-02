package service

import (
	"encoding/json"
	"github.com/Szzx123/depot_reparti/model/message"
	"github.com/Szzx123/depot_reparti/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

var ConnSnap *websocket.Conn

func Snapshot_Handler(c *gin.Context) {
	var err error
	ConnSnap, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Upgrade WebSocket Failure: %s", err)
		return
	}
	go Snapshot_Receive_Handler(c)
}

func Snapshot_Receive_Handler(c *gin.Context) {
	defer ConnSnap.Close()
	for {
		// Lire le message
		_, message_json, err := ConnSnap.ReadMessage()
		if err != nil {
			log.Printf("Read Message Failure: %s", err)
			break
		}

		// Stockage temporaire du contenu du message
		message_content := make([]byte, len(message_json))
		copy(message_content, message_json)

		// Analyse de JSON
		var msg message.SnapshotMessage
		if err := json.Unmarshal(message_content, &msg); err != nil {
			log.Printf("Analyse JSON Failure: %s", err)
			continue
		}

		// demande Snapshot
		utils.Msg_send(utils.Msg_format("receiver", "C"+msg.Site[1:]) + utils.Msg_format("type", "demandeSnapshot"))
		// Envoyer un message
		if err := ConnSnap.WriteMessage(websocket.TextMessage, []byte("demandeSnapshot已收到")); err != nil {
			log.Printf("Send Message Failure: %s", err)
			break
		}

	}
}

func Snapshot_Send_Handler(msgSnapshot message.SnapshotMessage) {
	// Marshal the Person struct to a JSON byte slice
	jsonSnapshot, err := json.Marshal(msgSnapshot)
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}

	// Print the JSON string
	log.Println(string(jsonSnapshot))

	if ConnSnap != nil {
		err = ConnSnap.WriteMessage(websocket.TextMessage, jsonSnapshot)
		if err != nil {
			log.Fatal(err)
		}
	}
}
