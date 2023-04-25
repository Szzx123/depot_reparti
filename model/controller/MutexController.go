package controller

import (
	"math"
	"time"

	"github.com/Szzx123/depot_reparti/model/message"
)

type MutexController struct {
	num     int                    //number of site
	tab     []message.MutexMessage //register the latest statues of all sites
	horloge int                    //horloge local
	channel chan string            //channel to communicate with basic application
}

func New_MutexController() *MutexController {
	channel := make(chan string)
	tab := make([]message.MutexMessage, 5) //怎么样知道初始有多少site？
	return &MutexController{
		num:     1, //根据全局变量来？
		horloge: 0,
		tab:     tab,
		channel: channel,
	}
}

// Réception d’une demande de section critique ou de fin de section critique de l’application de base
func (mc *MutexController) Receive_RequestSC() {
	for {
		select {
		case instruction := <-mc.channel:
			if instruction == "demandeSC" {
				mc.horloge += 1
				msg := mm.New_MutexMessage(mc.horloge, 1)
				mc.tab[mc.num] = *msg
				// envoyer( [requête] hi ) à tous les autres sites
			} else if instruction == "finSC" {
				mc.horloge += 1
				msg := mm.New_MutexMessage(mc.horloge, 0)
				mc.tab[mc.num] = *msg
				// envoyer( [libération] hi ) à tous les autres sites.
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Réception d’un message de type requête
// Réception d’un message de type libération
// Réception d’un message de type accusé
func (mc *MutexController) ExtMessage_Handler() {
	for {
		// recevoir( [type_message ] h ) de Sj
		if type_message == "requête" {
			mc.horloge = math.MaxInt(mc.horloge, ext_horloge) + 1
			mc.tab[ext_num] = *mm.New_MutexMessage(mc.horloge, 1)
			// envoyer( [accusé] hi ) à Sj
			mc.Send_StartSC(ext_num)
		} else if type_message == "libération" {
			mc.horloge = math.MaxInt(mc.horloge, ext_horloge) + 1
			mc.tab[ext_num] = *mm.New_MutexMessage(mc.horloge, 0)
			mc.Send_StartSC(ext_num)
		} else if type_message == "accusé" {
			mc.horloge = math.MaxInt(mc.horloge, ext_horloge) + 1
			if mc.tab[ext_num].typeMessage != 1 {
				mc.tab[ext_num] = *mm.New_MutexMessage(mc.horloge, 2)
			}
			mc.Send_StartSC(ext_num)
		}
	}
}

// L’arrivée du message pourrait permettre de satisfaire une éventuelle demande de Si.
func (mc *MutexController) Send_StartSC(ext_num int) {
	if mc.tab[mc.num] == 1 {
		for k := range mc.tab {
			if k != mc.num && utils.compare_timestamp(mc.tab[mc.num], mc.num, mc.tab[ext_num], ext_num) {
				break
			}
		}
		mc.channel <- "debutSC" // envoyer( [débutSC] ) à l’application de base
	}
}
