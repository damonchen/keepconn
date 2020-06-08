package master

import (
	"github.com/damonchen/keepconn/internal/message"
)

// // MinionChannel minion channel
// type MinionChannel struct {
// 	ID      string
// 	channel socks.Channel
// }

// Queue queue
type Queue interface {
	// Push message to back
	Push(message.Message) error
	// Pop message
	Pop() (message.Message, error)
}

// MessagePump message pump
type MessagePump struct {
	topic      string
	subscriber Subscriber

	messageChan chan message.Message
}

// Publish publish
func (m *MessagePump) Publish(msg message.Message) error {
	return nil
}

// pump pump message
func (m *MessagePump) pump() {
	return
}
