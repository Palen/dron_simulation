package subscribers

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/Palen/drone_simulation/pkg/geo"
)

var trafficChoices = [...]string{"HEAVY", "LIGHT", "MODERATE"}

type Drone struct {
	id          uint64
	checkpoints []*geo.CheckPoint
	channel     *MessageChannel
	lastTime    *time.Time
	lastCoord   *geo.Coord
	exit        chan *sync.WaitGroup
}

func (d *Drone) Send(message *Message) {
	*d.channel <- message

}

func (d *Drone) Exit(waiter *sync.WaitGroup) {
	d.exit <- waiter
}

func (d *Drone) Move(coords *geo.Coord, t *time.Time) {
	speed := 0.0
	if d.lastCoord == nil {
		d.lastCoord = coords
		d.lastTime = t
	} else {
		distanceFromLastPointM := d.lastCoord.Distance(coords)
		timeElapsed := float64(t.Unix()) - float64(d.lastTime.Unix())
		speed = distanceFromLastPointM / timeElapsed
		log.Println(timeElapsed, distanceFromLastPointM, speed)
		time.Sleep(time.Duration(timeElapsed/10000) * time.Second)
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

func (d *Drone) Subscribe() {
	for {
		select {
		case droneMessage := <-*d.channel:
			d.Move(&droneMessage.Coord, &droneMessage.Time)
		case waiter := <-d.exit:
			waiter.Done()
			return
		}
	}
}

func NewDrone(checkpts []*geo.CheckPoint, maxSize int, id uint64) *Drone {
	channel := make(MessageChannel, maxSize)
	quit := make(chan *sync.WaitGroup, 1)
	drone := Drone{checkpoints: checkpts, channel: &channel, id: id, lastTime: nil, exit: quit}
	return &drone
}
