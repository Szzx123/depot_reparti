package client

import (
	// "github.com/Szzx123/depot_reparti/model/controller"
	"fmt"
	"time"

	"github.com/Szzx123/depot_reparti/model/message"
)

type Client struct {
	Num             string
	Channel_message chan message.Message
	// mutexController *controller.MutexController
}

func New_Client(num string) *Client {
	return &Client{
		Num:             num,
		Channel_message: make(chan message.Message),
		// mutexController: controller.New_MutexController(num),
	}
}

func (client *Client) Run() {
	go client.Message_Handler()
}

func (client *Client) Message_Handler() {
	fmt.Println("Start of message_handler ", client.Num)
	for {
		select {
		case msg := <-client.Channel_message:
			fmt.Println("Message received", msg)
			//发送demandeSC给ctl
			//等待mutexctl的允许后进行操作
			//流程结束后发送finSC给ctl
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
