package controller

import (
	"time"

	"github.com/Szzx123/depot_reparti/model/message"
)

var (
	N int = 5
)

type ISnapshotController interface {
}
type SnapshotController struct {
	num           int //number of site
	bilan         int //Nombre d’émissions moins nombre de réception
	color         message.Color
	vectorHorloge []int                        //horloge vectorielle local
	channel       chan message.SnapshotMessage //channel to communicate with basic application
	initiator     bool
	nbStateToWait int           // États devant être reçus.
	nbMsgToWait   int           // Messages devant être reçus
	globalState   map[int][]int // key:number of site, val: horloge vectorielle
}

func New_SnapshotController(num int) *SnapshotController {
	channel := make(chan message.SnapshotMessage)
	global_state := make(map[int][]int)
	return &SnapshotController{
		num:           num,
		bilan:         0,
		color:         0,
		vectorHorloge: []int{},
		channel:       channel,
		initiator:     false,
		nbStateToWait: 0,
		nbMsgToWait:   0,
		globalState:   global_state,
	}
}

// Début de l’instantané
func (ss *SnapshotController) StartSnapshot() {
	ss.color = 1
	ss.initiator = true
	ss.globalState[ss.num] = ss.vectorHorloge
	ss.nbMsgToWait = ss.bilan
	ss.nbStateToWait = N - 1
}

// Envoi d’un message m de l’application de base :
func (ss *SnapshotController) Send_Message() {
	// envoyer( m, couleuri )
	ss.bilan += 1
}

// Réception d’un message de l’application de base :
func (ss *SnapshotController) BaseMessage_Handler() {
	for {
		select {
		case msg := <-ss.channel:
			ss.bilan -= 1
			// Première réception d’un message rouge. Si prend son instantané.
			if msg.Get_Color() == 1 && ss.color == 0 {
				ss.color = 1
				ss.globalState[ss.num] = ss.vectorHorloge
				// envoyer( [état] EG, bilani ) sur l’anneau
			}
			// Réception postclic d’un message envoyé préclic.
			if msg.Get_Color() == 0 && ss.color == 1 {
				// envoyer( [prépost] m ) sur l’anneau
			}
			// TODO: traiter le message m
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Réception d’un message de type [état]
// Réception d’un message de type [prépost]
func (ss *SnapshotController) ExternalMessage_Handler() {
	for {
		// recevoir( [état] EG, bilan ) /* Réception d’un état local et d’un bilan. */
		// recevoir( [prépost] m )
		if type_message == "état" {
			if ss.initiator {
				ss.globalState[ext_num] = ext_global_state
				ss.nbMsgToWait += ext_bilan
				ss.nbStateToWait -= 1
				if ss.nbStateToWait == 0 && ss.nbMsgToWait == 0 {
					return // Fin de l’algorithme
				}
			} else {
				// envoyer( [état] EG, bilan ) sur l’anneau
			}
		} else if type_message == "prépost" {
			if ss.initiator {
				ss.nbMsgToWait -= 1
				// EGi ← EGi ∪ m
				if ss.nbStateToWait == 0 && ss.nbMsgToWait == 0 {
					return // Fin de l’algorithme
				}
			} else {
				// envoyer( [état] EG, bilan ) sur l’anneau
			}
		}
	}
}
