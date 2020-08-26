package subscribers

import (
	"log"

	"github.com/Palen/drone_simulation/pkg/checkpoints"
)

type Drone struct {
	id          uint64
	checkpoints []*checkpoints.CheckPoint
	channel     *MessageChannel
}

func (d *Drone) Move(coords *checkpoints.Coord) {
	for _, checkpoint := range d.checkpoints {
		distance := checkpoint.Distance(coords)
		if distance < 350 {
			log.Println("In ", checkpoint.Name, distance, d.id)
		} else {
		}
	}
}

func (d *Drone) Send(message *Message) {
	*d.channel <- message

}

func (d *Drone) Subscribe() {
	for {
		select {
		case droneMessage := <-*d.channel:
			if h, m, _ := droneMessage.Time.Clock(); h >= 8 && m >= 10 {
				log.Println("Time is now: ", droneMessage.Time.String())
				return
			}
			d.Move(&droneMessage.Coord)
		}
	}
}

func NewDrone(checkpts []*checkpoints.CheckPoint, maxSize int, id uint64) *Drone {
	channel := make(MessageChannel, maxSize)
	drone := Drone{checkpoints: checkpts, channel: &channel, id: id}
	return &drone
}
