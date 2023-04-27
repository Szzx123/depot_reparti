package controller

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Szzx123/depot_reparti/model/message"
	"github.com/Szzx123/depot_reparti/utils"
)

type Controller struct {
	num     string                          //number of site
	tab     map[string]message.MutexMessage //register the latest statues of all sites
	horloge int                             //horloge local
}

func New_Controller(num string) *Controller {
	tab := make(map[string]message.MutexMessage)
	return &Controller{
		num:     num,
		horloge: 0,
		tab:     tab,
	}
}

func (ctl *Controller) Run() {
	go ctl.Message_Interceptor()
}

// Réception d’une demande de section critique ou de fin de section critique de l’application de base
func (ctl *Controller) Message_Interceptor() {
	var rcv_msg string
	l := log.New(os.Stderr, "", 0)
	// l.Printf(string(ctl.num))
	var receiver, sender string
	var msg_type int
	var logical_time int
	for {
		fmt.Scanln(&rcv_msg)
		l.Printf("controller %s received message: %s", ctl.num, rcv_msg)
		// mutex.Lock()
		tab_allkeyval := strings.Split(rcv_msg[1:], rcv_msg[0:1])
		for _, key_val := range tab_allkeyval {
			tab_keyval := strings.Split(key_val[1:], key_val[0:1])
			if tab_keyval[0] == "receiver" {
				receiver = tab_keyval[1]
				if receiver != "All" || receiver != ctl.num {
					fmt.Println(rcv_msg)
					break
				}
			} else if tab_keyval[0] == "type" {
				switch tab_keyval[1] {
				case "request":
					msg_type = 1
				case "release":
					msg_type = 0
				case "ack":
					msg_type = 2
				case "demandeSC":
					msg_type = 3
				case "finSC":
					msg_type = 4
				default:
					l.Println("Invalid message type. Please try again.")
				}
			} else if tab_keyval[0] == "sender" {
				sender = tab_keyval[1]
			} else if tab_keyval[0] == "hlg" {
				logical_time, _ = strconv.Atoi(tab_keyval[1])
			}
		}
		msg_to_handle := message.New_MutexMessage(sender, logical_time, message.TypeMessage(msg_type))
		ctl.Message_Handler(msg_to_handle)
	}
}

// Réception d’un message de type requête
// Réception d’un message de type libération
// Réception d’un message de type accusé
func (ctl *Controller) Message_Handler(msg *message.MutexMessage) {
	ext_num := msg.Get_Site()
	l := log.New(os.Stderr, "", 0)
	switch msg.Get_typeMessage() {
	case "demandeSC":
		ctl.horloge += 1
		msg := message.New_MutexMessage(ctl.num, ctl.horloge, 1)
		ctl.tab[ctl.num] = *msg
		// envoyer( [requête] hi ) à tous les autres sites
		fmt.Printf("/=receiver=All/=type=request/=sender=%s/=hlg=%d\n", ctl.num, ctl.horloge)
		l.Println(ctl.tab)
	case "finSC":
		ctl.horloge += 1
		msg := message.New_MutexMessage(ctl.num, ctl.horloge, 0)
		ctl.tab[ctl.num] = *msg
		// envoyer( [libération] hi ) à tous les autres sites.
		fmt.Printf("/=receiver=All/=type=release/=sender=%s/=hlg=%d\n", ctl.num, ctl.horloge)
		l.Println(ctl.tab)
	case "request":
		ctl.horloge = utils.Recaler(ctl.horloge, msg.Get_Horloge())
		ctl.tab[ext_num] = *msg
		// envoyer( [accusé] hi ) à Sj
		fmt.Printf("/=receiver=%s/=type=ack/=sender=%s/=hlg=%d\n", ext_num, ctl.num, ctl.horloge)
		ctl.Send_StartSC(ext_num)
		l.Println(ctl.tab) // test
	case "release":
		ctl.horloge = utils.Recaler(ctl.horloge, msg.Get_Horloge())
		ctl.tab[ext_num] = *msg
		ctl.Send_StartSC(ext_num)
		l.Println(ctl.tab) // test
	case "ack":
		ctl.horloge = utils.Recaler(ctl.horloge, msg.Get_Horloge())
		if ctl.tab[ext_num].Get_typeMessage() != "requête" {
			ctl.tab[ext_num] = *message.New_MutexMessage(ctl.num, ctl.horloge, 2)
		}
		ctl.Send_StartSC(ext_num)
		l.Println(ctl.tab) // test
	}
}

// L’arrivée du message pourrait permettre de satisfaire une éventuelle demande de Si.
func (ctl *Controller) Send_StartSC(ext_num string) {
	ext_num_int, _ := strconv.Atoi(ext_num)
	num, _ := strconv.Atoi(ctl.num)
	if ctl.tab[ctl.num].Get_typeMessage() == "requête" {
		for k := range ctl.tab {
			if k != ctl.num && utils.Compare_Timestamp(ctl.tab[ctl.num].Get_Horloge(), num, ctl.tab[ext_num].Get_Horloge(), ext_num_int) {
				break
			}
		}
		fmt.Printf("/=receiver=%s/=type=debutSC/=sender=%s/=hlg=%d\n", ext_num, ctl.num, ctl.horloge) // envoyer( [débutSC] ) à l’application de base
	}
}
