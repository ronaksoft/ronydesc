package desc

type IMessage interface {
	msg()
}

type Message struct{}

func (m Message) msg() {}
