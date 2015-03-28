package dummy

import (
	"github.com/Townsita/townsita"
	"time"
)

type Dummy struct{}

func New() *Dummy {
	return &Dummy{}
}

func (d *Dummy) Init() {
}

func (d *Dummy) MustGetMessageTypes() []*townsita.MessageType {
	return []*townsita.MessageType{
		&townsita.MessageType{"type1", "Message Type 1"},
		&townsita.MessageType{"type2", "Message Type 2"},
	}
}

func (d *Dummy) MustGetMessageSubTypes(id string) []*townsita.MessageType {
	return []*townsita.MessageType{
		&townsita.MessageType{"type1", "Message Type 1"},
		&townsita.MessageType{"type2", "Message Type 2"},
	}
}

func (d *Dummy) GetMessageTypeById(id string) *townsita.MessageType {
	return &townsita.MessageType{"type1", "Message Type 1"}
}

func (d *Dummy) SaveMessage(message *townsita.Message, user *townsita.User) (string, error) {
	return "message1", nil
}

func (d *Dummy) GetMessageById(id string) (*townsita.Message, error) {
	return &townsita.Message{
		ID:     "message1",
		UserID: "user1",
		TypeID: "type1",

		Readers:    100,
		Completed:  30,
		TargetHash: "XoXoXo",
		Latitude:   59.35,
		Longitude:  17.9167,
		Radius:     20,

		Headline: "Message Headline",
		Body:     "Message Body",
		Photo:    "https://i.imgur.com/u5vVrjT.gif",

		Status:    townsita.MessageDraft,
		CreatedAd: time.Now(),
		UpdatedAd: time.Now(),
	}, nil
}

func (d *Dummy) LoginUser(*townsita.User) (string, error) {
	return "user1", nil
}
