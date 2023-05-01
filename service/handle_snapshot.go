package service

import (
	"encoding/json"
	"github.com/Szzx123/depot_reparti/model/message"
	"github.com/Szzx123/depot_reparti/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

//var upgrader = websocket.Upgrader{
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}

var ConnSnap *websocket.Conn

func Snapshot_Handler(c *gin.Context) {
	var err error
	ConnSnap, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("升级 WebSocket 失败: %s", err)
		return
	}
	//defer Conn_Snap.Close()
	go Snapshot_Receive_Handler(c)
	//go Snapshot_Send_Handler()

}

func Snapshot_Receive_Handler(c *gin.Context) {
	defer ConnSnap.Close()
	for {
		// 读取消息
		//_, message_json, err := conn.ReadMessage()
		_, message_json, err := ConnSnap.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %s", err)
			break
		}

		// 暂存消息内容
		message_content := make([]byte, len(message_json))
		copy(message_content, message_json)

		// 解析 JSON
		var msg message.SnapshotMessage
		if err := json.Unmarshal(message_content, &msg); err != nil {
			log.Printf("解析 JSON 失败: %s", err)
			continue
		}

		// demande Snapshot
		utils.Msg_send(utils.Msg_format("receiver", "C"+msg.Site[1:]) + utils.Msg_format("type", "demandeSnapshot"))
		// 发送消息
		if err := ConnSnap.WriteMessage(websocket.TextMessage, []byte("demandeSnapshot已收到")); err != nil {
			log.Printf("发送消息失败: %s", err)
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

	//l.Println(message_snapshot)
	if ConnSnap != nil {
		err = ConnSnap.WriteMessage(websocket.TextMessage, jsonSnapshot)
		if err != nil {
			log.Fatal(err)
		}
	}
}
