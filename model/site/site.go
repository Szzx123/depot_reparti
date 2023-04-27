package site

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Szzx123/depot_reparti/global/depot"
	"github.com/Szzx123/depot_reparti/model/message"
	"github.com/Szzx123/depot_reparti/utils"
	"github.com/Szzx123/depot_reparti/utils/timestamp"
)

var (
	Site_1                  = New_Site("8080")
	Site_2                  = New_Site("8081")
	Site_3                  = New_Site("8082")
	Sites  map[string]*Site = map[string]*Site{
		"8080": Site_1,
		"8081": Site_2,
		"8082": Site_3,
	}
)

type Site struct {
	Num             string
	Channel_message chan message.Message
	Lock            *Lock
}

type Lock struct {
	num          string                          //number of site
	tab          map[string]message.MutexMessage //register the latest statues of all sites
	horloge      int                             //horloge local
	Channel_base chan string                     //channel to communicate with basic application
	Channel_ctl  chan message.MutexMessage       // channel to communicate with other ctls
}

func New_Site(num string) *Site {
	return &Site{
		Num:             num,
		Channel_message: make(chan message.Message),
		Lock:            New_Lock(num),
	}
}

func New_Lock(num string) *Lock {
	channel_base := make(chan string)
	channel_ctl := make(chan message.MutexMessage, 6)
	tab := make(map[string]message.MutexMessage)
	return &Lock{
		num:          num,
		horloge:      0,
		tab:          tab,
		Channel_base: channel_base,
		Channel_ctl:  channel_ctl,
	}
}

func (site *Site) Run() {
	go site.Message_Handler()
	go site.Lock.Receive_RequestSC()
	go site.Lock.ExtMessage_Handler()
}

func (site *Site) Message_Handler() {
	fmt.Println("Start of message_handler ", site.Num)
	for {
		select {
		case msg := <-site.Channel_message:
			go func() {
				site.Lock.Channel_base <- "demandeSC"
				reply := <-site.Lock.Channel_base
				if reply == "debutSC" {
					//等待mutexctl的允许后进行库存操作
					quantity, _ := strconv.Atoi(msg.Quantity)
					if msg.Type == "in" {
						depot.Depot.Cargo_IN(msg.Cargo, quantity)
					} else if msg.Type == "out" {
						depot.Depot.Cargo_OUT(msg.Cargo, quantity)
					}
				}
				//流程结束后发送finSC给ctl
				site.Lock.Channel_base <- "finSC"
			}()
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Réception d’une demande de section critique ou de fin de section critique de l’application de base
func (mc *Lock) Receive_RequestSC() {
	for {
		select {
		case instruction := <-mc.Channel_base:
			if instruction == "demandeSC" {
				fmt.Println(mc.num, "收到demandeSC")
				mc.horloge += 1
				msg := message.New_MutexMessage(mc.num, mc.horloge, 1)
				mc.tab[mc.num] = *msg
				// envoyer( [requête] hi ) à tous les autres sites
				mc.Diffusion(*msg)
				fmt.Println(mc.tab)
			} else if instruction == "finSC" {
				fmt.Println(mc.num, "收到finSC")
				mc.horloge += 1
				msg := message.New_MutexMessage(mc.num, mc.horloge, 0)
				mc.tab[mc.num] = *msg
				// envoyer( [libération] hi ) à tous les autres sites.
				mc.Diffusion(*msg)
				fmt.Println(mc.tab)
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Réception d’un message de type requête
// Réception d’un message de type libération
// Réception d’un message de type accusé
func (mc *Lock) ExtMessage_Handler() {
	for {
		select {
		case msg := <-mc.Channel_ctl:
			ext_num := msg.Get_Site()
			if msg.Get_typeMessage() == "requête" {
				mc.horloge = utils.Recaler(mc.horloge, msg.Get_Horloge())
				mc.tab[ext_num] = *message.New_MutexMessage(mc.num, mc.horloge, 1)
				// envoyer( [accusé] hi ) à Sj
				mc.Send(ext_num, *message.New_MutexMessage(mc.num, mc.horloge, 2))
				mc.Send_StartSC(ext_num)
				fmt.Println(mc.tab)
			} else if msg.Get_typeMessage() == "libération" {
				mc.horloge = utils.Recaler(mc.horloge, msg.Get_Horloge())
				mc.tab[ext_num] = *message.New_MutexMessage(mc.num, mc.horloge, 0)
				mc.Send_StartSC(ext_num)
				fmt.Println(mc.tab)
			} else if msg.Get_typeMessage() == "accusé" {
				mc.horloge = utils.Recaler(mc.horloge, msg.Get_Horloge())
				if mc.tab[ext_num].Get_typeMessage() != "requête" {
					mc.tab[ext_num] = *message.New_MutexMessage(mc.num, mc.horloge, 2)
				}
				mc.Send_StartSC(ext_num)
				fmt.Println(mc.tab)
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}

	}
}

// L’arrivée du message pourrait permettre de satisfaire une éventuelle demande de Si.
func (mc *Lock) Send_StartSC(ext_num string) {
	ext_num_int, _ := strconv.Atoi(ext_num)
	num, _ := strconv.Atoi(mc.num)
	if mc.tab[mc.num].Get_typeMessage() == "requête" {
		for k := range mc.tab {
			if k != mc.num && timestamp.Compare_Timestamp(mc.tab[mc.num].Get_Horloge(), num, mc.tab[ext_num].Get_Horloge(), ext_num_int) {
				break
			}
		}
		mc.Channel_base <- "debutSC" // envoyer( [débutSC] ) à l’application de base
	}
}

func (mc *Lock) Diffusion(msg message.MutexMessage) {
	for site := range Sites {
		if site != mc.num {
			Sites[site].Lock.Channel_ctl <- msg
		}
	}
}

func (mc *Lock) Send(target_site string, msg message.MutexMessage) {
	Sites[target_site].Lock.Channel_ctl <- msg
}
