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

func (d *Dummy) MustGetMessageSubTypes(id string) []*townsita.MessageType {
	return []*townsita.MessageType{
		&townsita.MessageType{"1", "Message Type 1"},
		&townsita.MessageType{"2", "Message Type 2"},
	}
}

func (d *Dummy) GetMessageTypeById(id string) *townsita.MessageType {
	return &townsita.MessageType{"1", "Message Type 1"}
}

func (d *Dummy) SaveMessage(message *townsita.Message, user *townsita.User) (string, error) {
	return "1", nil
}
