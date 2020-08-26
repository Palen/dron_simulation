package dispatcher

import (
	"log"

	"github.com/Palen/drone_simulation/pkg/subscribers"
)

type DispatcherChan chan string

type Dispatcher struct {
	channel *DispatcherChan
}

func (d *Dispatcher) Start(subs subscribers.Subscribers) {
	for {
		select {
		case line := <-*d.channel:
			droneMessage, err := subscribers.NewMessage(line)
			if err != nil {
				log.Println("Invalid message with err:", err)
			} else {
				if subscriber, ok := subs[droneMessage.Id]; ok {
					subscriber.Send(droneMessage)
				}
			}
		}
	}
}

func New(dispatcherChan *DispatcherChan) *Dispatcher {
	dispatcher := Dispatcher{dispatcherChan}
	return &dispatcher
}
