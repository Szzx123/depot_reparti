package site

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	// Site_1                  = New_Site("8080")
	// Site_2                  = New_Site("8081")
	// Site_3                  = New_Site("8082")
	// Sites  map[string]*Site = map[string]*Site{
	// 	"8080": Site_1,
	// 	"8081": Site_2,
	// 	"8082": Site_3,
	// }
	mutex = &sync.Mutex{}
)

type Site struct {
	Num string
}

func New_Site(num string) *Site {
	return &Site{
		Num: num,
	}
}

func (site *Site) Run() {
	go site.Message_Interceptor()
}

func (site *Site) Message_Interceptor() {
	var rcv_msg, cargo, msg_type, operation, sender string
	var quantity int
	var to_operate_cargo bool = false
	l := log.New(os.Stderr, "", 0)
	l.Println("Start of message_handler ", site.Num)
	for {
		fmt.Scanln(&rcv_msg)
		// mutex.Lock()
		l.Printf("site %s received message: %s", site.Num, rcv_msg)
		tab_allkeyval := strings.Split(rcv_msg[1:], rcv_msg[0:1])
		for _, key_val := range tab_allkeyval {
			tab_keyval := strings.Split(key_val[1:], key_val[0:1])
			if tab_keyval[0] == "receiver" {
				receiver := tab_keyval[1]
				if receiver != site.Num {
					break
				}
			} else if tab_keyval[0] == "type" && tab_keyval[1] == "débutSC" {
				//允许后进行库存操作
				to_operate_cargo = true
			} else if tab_keyval[0] == "cargo" {
				cargo = tab_keyval[1]
			} else if tab_keyval[0] == "operation" {
				operation = tab_keyval[1]
			} else if tab_keyval[0] == "quantity" {
				quantity, _ = strconv.Atoi(tab_keyval[1])
			} else if tab_keyval[0] == "sender" {
				sender = tab_keyval[1]
			}
		}
		if to_operate_cargo {
			l.Println(cargo, msg_type, operation, quantity, sender) ////////
		}
	}
}
