package dispatcher

import (
	"log"
	"sync"

	"github.com/Palen/drone_simulation/pkg/subscribers"
)

type DispatcherChan chan string

type Dispatcher struct {
	channel *DispatcherChan
}

func (d *Dispatcher) Start(subs subscribers.Subscribers, waiter *sync.WaitGroup) {

	for {
		select {
		case line := <-*d.channel:
			droneMessage, err := subscribers.NewMessage(line)
			if err != nil {
				log.Println("Invalid message with err:", err)
			}
			if h, m, _ := droneMessage.Time.Clock(); h >= 8 && m >= 10 {
				// If time >= 8:10 kill subscriber goroutine
				if subscriber, ok := subs[droneMessage.Id]; ok {
					subscriber.Exit(waiter)
					delete(subs, droneMessage.Id)
				}
			}
			if subscriber, ok := subs[droneMessage.Id]; ok {
				subscriber.Send(droneMessage)
			}
		}
	}
}

func New(dispatcherChan *DispatcherChan) *Dispatcher {
	dispatcher := Dispatcher{dispatcherChan}
	return &dispatcher
}
