package subscribers

import "sync"

type Subscriber interface {
	Send(m *Message)
	Exit(w *sync.WaitGroup)
}

type Subscribers map[uint64]Subscriber
