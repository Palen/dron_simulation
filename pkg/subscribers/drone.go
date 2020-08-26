package subscribers

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Palen/drone_simulation/pkg/checkpoints"
)

var trafficChoices = [...]string{"HEAVY", "LIGHT", "MODERATE"}

type Drone struct {
	id          uint64
	checkpoints []*checkpoints.CheckPoint
	channel     *MessageChannel
	lastTime    *time.Time
	lastCoord   *checkpoints.Coord
}

func (d *Drone) Move(coords *checkpoints.Coord, t *time.Time) {
	speed := 0.0
	if d.lastCoord == nil {
		d.lastCoord = coords
		d.lastTime = t
	} else {
		distanceFromLastPointM := d.lastCoord.Distance(coords)
		timeElapsed := float64(t.Unix()) - float64(d.lastTime.Unix())
		speed = distanceFromLastPointM / timeElapsed
		time.Sleep(time.Duration(timeElapsed) * time.Second)
	}
	for _, checkpoint := range d.checkpoints {
		distance := checkpoint.Coord.Distance(coords)
		if distance < 350 {
			conditions := trafficChoices[rand.Intn(len(trafficChoices))]
			log.Println(fmt.Sprintf("Traffic Report: ID: %d Time: %s Speed: %.2fm/s Traffic: %s",
				d.id, t, speed, conditions))
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
			d.Move(&droneMessage.Coord, &droneMessage.Time)
		}
	}
}

func NewDrone(checkpts []*checkpoints.CheckPoint, maxSize int, id uint64) *Drone {
	channel := make(MessageChannel, maxSize)
	drone := Drone{checkpoints: checkpts, channel: &channel, id: id, lastTime: nil}
	return &drone
}
