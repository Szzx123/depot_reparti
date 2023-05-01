package controller

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/Szzx123/depot_reparti/model/message"
	"github.com/Szzx123/depot_reparti/utils"
)

var (
	mutex = &sync.Mutex{}
)

type Controller struct {
	num         string                          //number of site
	tab         map[string]message.MutexMessage //register the latest statues of all sites
	horloge     int                             //horloge local
	ok          bool
	horloge_vec []int
	color       int
	snapshot    string
}

func New_Controller(num string) *Controller {
	tab := make(map[string]message.MutexMessage)
	msg_1 := message.New_MutexMessage("C1", 1, 0, "", 0, "", 0, 0, 0, 1, 1, 1)
	msg_2 := message.New_MutexMessage("C2", 1, 0, "", 0, "", 0, 0, 0, 1, 1, 1)
	msg_3 := message.New_MutexMessage("C3", 1, 0, "", 0, "", 0, 0, 0, 1, 1, 1)
	tab["C1"] = *msg_1
	tab["C2"] = *msg_2
	tab["C3"] = *msg_3
	horloge_vec := make([]int, 3)
	for i := 0; i < len(horloge_vec); i++ {
		horloge_vec[i] = 1
	}
	return &Controller{
		num:         num,
		horloge:     1,
		tab:         tab,
		horloge_vec: horloge_vec,
		color:       0,
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
	for {
		var receiver, sender, cargo, operation string
		var quantity, stock_A, stock_B, stock_C int
		var logical_time int
		var msg_type int
		fmt.Scanln(&rcv_msg)
		mutex.Lock()

		// treat champ sender
		sender = utils.Findval(rcv_msg, "sender")

		// treat champs receiver
		receiver = utils.Findval(rcv_msg, "receiver")
		if receiver == "" || sender == ctl.num {
			mutex.Unlock()
			continue
		} else if receiver != ctl.num {
			utils.Msg_send(rcv_msg)
			mutex.Unlock()
			continue
		}
		l.Printf("controller %s received message: %s", ctl.num, rcv_msg)

		// treat champs type
		msg_type_string := utils.Findval(rcv_msg, "type")
		if msg_type_string != "" {
			switch msg_type_string {
			case "request":
				msg_type = 1
			case "release":
				msg_type = 0
			case "ack":
				msg_type = 2
			case "demandeSC":
				msg_type = 3
			case "finSC":
				msg_type = 5
			case "demandeSnapshot":
				msg_type = 6
			case "finSnapshot":
				msg_type = 7
			default:
				l.Println("Unknown Message Type")
				mutex.Unlock()
				continue
			}
		}

		// treat champ horloge
		hlg := utils.Findval(rcv_msg, "horloge")
		if hlg != "" {
			logical_time, _ = strconv.Atoi(hlg)
		}

		// treat champ cargo
		cargo = utils.Findval(rcv_msg, "cargo")

		// treat champ operation
		operation = utils.Findval(rcv_msg, "operation")

		// treat champ quantity
		quantity_string := utils.Findval(rcv_msg, "quantity")
		if quantity_string != "" {
			quantity, _ = strconv.Atoi(quantity_string)
		}

		stock_A_string := utils.Findval(rcv_msg, "A")
		if stock_A_string != "" {
			stock_A, _ = strconv.Atoi(stock_A_string)
		}

		stock_B_string := utils.Findval(rcv_msg, "B")
		if stock_B_string != "" {
			stock_B, _ = strconv.Atoi(stock_B_string)
		}

		stock_C_string := utils.Findval(rcv_msg, "C")
		if stock_C_string != "" {
			stock_C, _ = strconv.Atoi(stock_C_string)
		}

		msg_to_handle := message.New_MutexMessage(sender, logical_time, message.TypeMessage(msg_type), cargo, quantity, operation, stock_A, stock_B, stock_C, ctl.horloge_vec[0], ctl.horloge_vec[1], ctl.horloge_vec[2])
		time.Sleep(1 * time.Second)
		ctl.Message_Handler(msg_to_handle)
		mutex.Unlock()
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
		// Mettre à jour l'horloge vectorielle
		num, err := strconv.Atoi(ctl.num[1:])
		if err != nil {
			// Handle error
		}
		ctl.horloge_vec[num-1] += 1

		new_msg := message.New_MutexMessage(ctl.num, ctl.horloge, 1, msg.Cargo, msg.Quantity, msg.Operation, 0, 0, 0, ctl.horloge_vec[0], ctl.horloge_vec[1], ctl.horloge_vec[2])
		ctl.tab[ctl.num] = *new_msg

		ctl.snapshot = ctl.snapshot + "opération : " + msg.Cargo + " " + strconv.Itoa(msg.Quantity) + " " + msg.Operation

		// envoyer( [requête] hi ) à tous les autres sites
		for i := 1; i <= 3; i++ {
			if strconv.Itoa(i) != ctl.num[1:] {
				utils.Msg_send(utils.Msg_format("receiver", "C"+strconv.Itoa(i)) + utils.Msg_format("type", "request") + utils.Msg_format("sender", ctl.num) + utils.Msg_format("horloge", strconv.Itoa(ctl.horloge)))
			}
		}
		ctl.ok = true
		l.Println(ctl.num, ": ", ctl.tab) // test
	case "finSC":
		ctl.horloge += 1
		// Mettre à jour l'horloge vectorielle
		num, err := strconv.Atoi(ctl.num[1:])
		if err != nil {
			// Handle error
		}
		ctl.horloge_vec[num-1] += 1

		stock_A := msg.Stock_A
		stock_B := msg.Stock_B
		stock_C := msg.Stock_C
		new_msg := message.New_MutexMessage(ctl.num, ctl.horloge, 0, "", 0, "", stock_A, stock_B, stock_C, ctl.horloge_vec[0], ctl.horloge_vec[1], ctl.horloge_vec[2])
		ctl.tab[ctl.num] = *new_msg

		ctl.snapshot = ctl.snapshot + ", horloge vectorielle [" + strconv.Itoa(ctl.horloge_vec[0]) + " " + strconv.Itoa(ctl.horloge_vec[1]) + " " + strconv.Itoa(ctl.horloge_vec[2]) + "]"

		// envoyer( [libération] hi ) à tous les autres sites.
		for i := 1; i <= 3; i++ {
			if strconv.Itoa(i) != ctl.num[1:] {
				utils.Msg_send(utils.Msg_format("receiver", "C"+strconv.Itoa(i)) + utils.Msg_format("type", "release") + utils.Msg_format("sender", ctl.num) + utils.Msg_format("horloge", strconv.Itoa(ctl.horloge)) + utils.Msg_format("A", strconv.Itoa(stock_A)) + utils.Msg_format("B", strconv.Itoa(stock_B)) + utils.Msg_format("C", strconv.Itoa(stock_C)))
			}
		}
		// utils.Msg_send(utils.Msg_format("receiver", "All") + utils.Msg_format("type", "release") + utils.Msg_format("sender", ctl.num) + utils.Msg_format("horloge", strconv.Itoa(ctl.horloge)) + utils.Msg_format("A", strconv.Itoa(stock_A)) + utils.Msg_format("B", strconv.Itoa(stock_B)) + utils.Msg_format("C", strconv.Itoa(stock_C)))
		l.Println(ctl.num, ": ", ctl.tab) // test
	case "request":
		ctl.horloge = utils.Recaler(ctl.horloge, msg.Get_Horloge())

		// Mettre à jour l'horloge vectorielle
		arr := []int{msg.H1, msg.H2, msg.H3}
		ctl.horloge_vec = utils.RecalerVec(ctl.horloge_vec, arr)
		num, err := strconv.Atoi(ctl.num[1:])
		if err != nil {
			// Handle error
		}
		ctl.horloge_vec[num-1] += 1

		ctl.tab[ext_num] = *msg
		// envoyer( [accusé] hi ) à Sj
		utils.Msg_send(utils.Msg_format("receiver", ext_num) + utils.Msg_format("type", "ack") + utils.Msg_format("sender", ctl.num) + utils.Msg_format("horloge", strconv.Itoa(ctl.horloge)))
		ctl.Send_StartSC()
		l.Println(ctl.num, ": ", ctl.tab) // test
	case "release":
		ctl.horloge = utils.Recaler(ctl.horloge, msg.Get_Horloge())

		// Mettre à jour l'horloge vectorielle
		arr := []int{msg.H1, msg.H2, msg.H3}
		ctl.horloge_vec = utils.RecalerVec(ctl.horloge_vec, arr)
		num, err := strconv.Atoi(ctl.num[1:])
		if err != nil {
			// Handle error
		}
		ctl.horloge_vec[num-1] += 1

		stock_A := msg.Stock_A
		stock_B := msg.Stock_B
		stock_C := msg.Stock_C
		ctl.tab[ext_num] = *msg
		utils.Msg_send(utils.Msg_format("receiver", "A"+ctl.num[1:]) + utils.Msg_format("type", "updateSC") + utils.Msg_format("sender", ctl.num) + utils.Msg_format("horloge", strconv.Itoa(ctl.horloge)) + utils.Msg_format("A", strconv.Itoa(stock_A)) + utils.Msg_format("B", strconv.Itoa(stock_B)) + utils.Msg_format("C", strconv.Itoa(stock_C)))
		ctl.Send_StartSC()
		l.Println(ctl.num, ": ", ctl.tab) // test
	case "ack":
		ctl.horloge = utils.Recaler(ctl.horloge, msg.Get_Horloge())

		// Mettre à jour l'horloge vectorielle
		arr := []int{msg.H1, msg.H2, msg.H3}
		ctl.horloge_vec = utils.RecalerVec(ctl.horloge_vec, arr)
		num, err := strconv.Atoi(ctl.num[1:])
		if err != nil {
			// Handle error
		}
		ctl.horloge_vec[num-1] += 1

		if ctl.tab[ext_num].Get_typeMessage() != "request" {
			ctl.tab[ext_num] = *msg
		}
		ctl.Send_StartSC()
		l.Println(ctl.num, ": ", ctl.tab) // test
	case "demandeSnapshot":
		if ctl.color == 0 { // blanc, traiter la demande de snapshot
			//if ext_num == ctl.num { // initiateur de snapshot
			//	initiateur =
			//}

			num, err := strconv.Atoi(ctl.num[1:])
			if err != nil {
				// Handle error
			}

			// envoyer message de demande de snapshot au site suivant
			utils.Msg_send(utils.Msg_format("receiver", "C"+strconv.Itoa(num%3+1)) + utils.Msg_format("type", "demandeSnapshot") + utils.Msg_format("sender", ctl.num) + utils.Msg_format("horloge", strconv.Itoa(ctl.horloge)))

			// generate snapshot
			horloge_snapshot := "[" + strconv.Itoa(ctl.horloge_vec[0]) + " " + strconv.Itoa(ctl.horloge_vec[1]) + " " + strconv.Itoa(ctl.horloge_vec[2]) + "]"
			utils.Msg_send(utils.Msg_format("receiver", "A"+ctl.num[1:]+utils.Msg_format("type", "generateSnapshot")+utils.Msg_format("sender", ctl.num)+utils.Msg_format("horloge_snapshot", horloge_snapshot)+utils.Msg_format("snapshot", ctl.snapshot)))

			// changer la couleur du site à rouge
			ctl.color = 1
		} else { // rouge, envoyer message de fin de snapshot

			num, err := strconv.Atoi(ctl.num[1:])
			if err != nil {
				// Handle error
			}
			utils.Msg_send(utils.Msg_format("receiver", "C"+strconv.Itoa(num%3+1)) + utils.Msg_format("type", "finSnapshot") + utils.Msg_format("sender", ctl.num))
		}
	case "finSnapshot":
		// change la couleur à blanc
		if ctl.color == 1 {
			num, err := strconv.Atoi(ctl.num[1:])
			if err != nil {
				// Handle error
			}
			utils.Msg_send(utils.Msg_format("receiver", "C"+strconv.Itoa(num%3+1)) + utils.Msg_format("type", "finSnapshot") + utils.Msg_format("sender", ctl.num))

			ctl.color = 0
		}
	}
}

// L’arrivée du message pourrait permettre de satisfaire une éventuelle demande de Si.
func (ctl *Controller) Send_StartSC() {
	if !ctl.ok {
		return
	}
	num, _ := strconv.Atoi(ctl.num[1:])
	// l := log.New(os.Stderr, "", 0)
	// l.Println(ctl.num, "尝试进入Section Critique")
	if ctl.tab[ctl.num].Get_typeMessage() == "request" {
		for k := range ctl.tab {
			ext_num, _ := strconv.Atoi(k[1:])
			// l.Println("本人时钟：", ctl.tab[ctl.num].Get_Horloge(), "本人站点号", num, "对比时钟", ctl.tab[k].Get_Horloge(), "对比站点号", ext_num)
			if k != ctl.num && !utils.Compare_Timestamp(ctl.tab[ctl.num].Get_Horloge(), num, ctl.tab[k].Get_Horloge(), ext_num) {
				// l.Println(ctl.num, "进入Section Critique失败")
				return
			}
		}
		// l.Println(ctl.num, "进入Section Critique成功")
		utils.Msg_send(utils.Msg_format("receiver", "A"+ctl.num[1:]) + utils.Msg_format("type", "débutSC") + utils.Msg_format("sender", ctl.num) + utils.Msg_format("horloge", strconv.Itoa(ctl.horloge)) + utils.Msg_format("cargo", ctl.tab[ctl.num].Cargo) + utils.Msg_format("operation", ctl.tab[ctl.num].Operation) + utils.Msg_format("quantity", strconv.Itoa(ctl.tab[ctl.num].Quantity)))
		ctl.ok = false
	}
}
