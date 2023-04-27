package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Szzx123/depot_reparti/model/message"
	"github.com/Szzx123/depot_reparti/model/site"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Cargo_Handler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("升级 WebSocket 失败: %s", err)
		return
	}
	defer conn.Close()
	for {
		// 读取消息
		_, message_json, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %s", err)
			break
		}

		// 暂存消息内容
		message_content := make([]byte, len(message_json))
		copy(message_content, message_json)

		// 解析 JSON
		var msg message.Message
		if err := json.Unmarshal(message_content, &msg); err != nil {
			log.Printf("解析 JSON 失败: %s", err)
			continue
		}

		if len(msg.Cargo) == 0 || len(msg.Type) == 0 || len(msg.Quantity) == 0 {
			log.Printf("Blank Message")
			continue
		}

		fmt.Printf("site: %s, cargo: %s, operation: %s, quantity: %s\n", msg.Site, msg.Cargo, msg.Type, msg.Quantity)

		// 发送消息
		if err := conn.WriteMessage(websocket.TextMessage, []byte("消息已收到")); err != nil {
			log.Printf("发送消息失败: %s", err)
			break
		}
		site.Sites[msg.Site].Channel_message <- msg
	}
}
