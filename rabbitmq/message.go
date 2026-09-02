package rabbitmq

import "github.com/google/uuid"

var (
	CMD_SUBFINDER  string = "subfinder"
	TYPE_SUBDOMAIN        = "subdomain"
)

type Message struct {
	UUID string   `json:"uuid" bson:"uuid"`
	CMD  string   `json:"cmd" bson:"cmd"`
	Args []string `json:"args" bson:"args"`
	Type string   `json:"type" bson:"type"`
}

func NewMessage(cmd string, msg_type string, args ...string) *Message {
	return &Message{UUID: uuid.NewString(), Type: msg_type, CMD: cmd, Args: args}

}
