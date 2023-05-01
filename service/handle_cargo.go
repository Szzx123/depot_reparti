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
		log.Printf("升级 WebSocket 失败: %s", err)
		return
	}
	//defer Conn_Snap.Close()
	go Cargo_Receive_Handler(c)
	//go Snapshot_Send_Handler()

}

func Cargo_Receive_Handler(c *gin.Context) {
	defer ConnCargo.Close()
	for {
		// 读取消息
		_, message_json, err := ConnCargo.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %s", err)
			break
		}

		// 暂存消息内容
		message_content := make([]byte, len(message_json))
		copy(message_content, message_json)

		// 解析 JSON
		var msg message.CargoMessage
		if err := json.Unmarshal(message_content, &msg); err != nil {
			log.Printf("解析 JSON 失败: %s", err)
			continue
		}

		if msg.TypeMessage == "operateCargo" {
			if len(msg.Cargo) == 0 || len(msg.Type) == 0 || len(msg.Quantity) == 0 {
				log.Printf("Blank Message")
				continue
			}

			utils.Msg_send(utils.Msg_format("receiver", "C"+msg.Site[1:]) + utils.Msg_format("type", "demandeSC") + utils.Msg_format("cargo", msg.Cargo) + utils.Msg_format("operation", msg.Type) + utils.Msg_format("quantity", msg.Quantity))
			// 发送消息
			//if err := ConnCargo.WriteMessage(websocket.TextMessage, []byte("消息已收到")); err != nil {
			//	log.Printf("发送消息失败: %s", err)
			//	break
			//}
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

	//l.Println(message_snapshot)
	if ConnCargo != nil {
		err = ConnCargo.WriteMessage(websocket.TextMessage, jsonCargo)
		if err != nil {
			log.Fatal(err)
		}
	}
}
