package dummy

import (
	"github.com/Townsita/townsita"
)

type Dummy struct{}

func New() *Dummy {
	return &Dummy{}
}

func (d *Dummy) Init() {
}

func (d *Dummy) MustGetMessageTypes() []*townsita.MessageType {
	return []*townsita.MessageType{
		&townsita.MessageType{"1", "Message Type 1"},
		&townsita.MessageType{"2", "Message Type 2"},
	}
}

func (d *Dummy) MustGetMessageSubTypes(messageType int) []*townsita.MessageType {
	return []*townsita.MessageType{
		&townsita.MessageType{"1", "Message Type 1"},
		&townsita.MessageType{"2", "Message Type 2"},
	}
}
