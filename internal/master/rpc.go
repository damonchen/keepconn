package master

import (
	"sync"

	"github.com/damonchen/keepconn/internal/message"
	"github.com/damonchen/keepconn/internal/socks"
)

// AgentStub agent stub
type AgentStub struct {
	ID      string
	session socks.Session
}

func (a *AgentStub) NewMessageChannel() *MessageChannel {
	return NewMessageChannel(a.ID)
}

type AgentChannel struct {
	messageChannel *MessageChannel
	agentStub      *AgentStub
}

func (a *AgentChannel) Publish(msg message.Message) {
	a.messageChannel.Publish(msg)
}

type AgentChannelManager struct {
	mutex    sync.Mutex
	channels map[string]*AgentChannel
}

func (m *AgentChannelManager) Get(id string) *AgentChannel {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.channels[id]
}

func (m *AgentChannelManager) Add(id string, channel *AgentChannel) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.channels[id] = channel
	return
}

// MessageChannel message channel
type MessageChannel struct {
	ID          string
	messageChan chan message.Message
}

func NewMessageChannel(ID string) *MessageChannel {
	msgChannel := &MessageChannel{
		ID:          ID,
		messageChan: make(chan message.Message, 1),
	}
	return msgChannel
}

func (m *MessageChannel) Publish(msg message.Message) {
	m.messageChan <- msg
}
