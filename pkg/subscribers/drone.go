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
	lastCoord   *geo.Coord
	exit        chan *sync.WaitGroup
	speed       float64
	perimeter   float64
}

func (d *Drone) Send(message *Message) {
	*d.channel <- message

}

func (d *Drone) Exit(waiter *sync.WaitGroup) {
	d.exit <- waiter
}

func (d *Drone) Subscribe() {
	for {
		select {
		case droneMessage := <-*d.channel:
			d.Move(droneMessage.Coord, &droneMessage.Time)
		case waiter := <-d.exit:
			waiter.Done()
			log.Println(fmt.Sprintf("Shutdown drone id: %d", d.id))
			return
		}
	}
}

func (d *Drone) Move(coords *geo.Coord, t *time.Time) {
	if d.lastCoord == nil {
		d.lastCoord = coords
	} else {
		// Moving drone
		distanceFromLastPointM := d.lastCoord.Distance(coords)
		log.Println(fmt.Sprintf("Drone ID: %d | distance from next point: %.2fm", d.id,
			distanceFromLastPointM))
		time.Sleep(time.Duration(distanceFromLastPointM/d.speed) * time.Second)
		d.lastCoord = coords
		for _, checkpoint := range d.checkpoints {
			distance := checkpoint.Coord.Distance(coords)
			if distance <= d.perimeter {
				conditions := trafficChoices[rand.Intn(len(trafficChoices))]
				log.Println(fmt.Sprintf("Traffic Report ID: %d | Time: %s | Speed: %.2fm/s | Traffic: %s",
					d.id, t, d.speed, conditions))
			}
		}
	}

}

func NewDrone(checkpts []*geo.CheckPoint, maxSize int, id uint64, speed float64, perimeter float64) *Drone {
	channel := make(MessageChannel, maxSize)
	quit := make(chan *sync.WaitGroup, 1)
	drone := Drone{checkpoints: checkpts, channel: &channel, id: id, exit: quit, speed: speed,
		perimeter: perimeter}
	return &drone
}
