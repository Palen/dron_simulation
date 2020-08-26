package subscribers

type Subscriber interface {
	Send(m *Message)
}

type Subscribers map[uint64]Subscriber
