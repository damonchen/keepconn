package Agent

import (
	"sync"

	"github.com/damonchen/keepconn/internal/agent/function"
	"github.com/damonchen/keepconn/internal/message"
	"github.com/damonchen/keepconn/internal/socks"
)

// Agent agent
type Agent struct {
	session    AgenetSession
	cmdChannel CommandChannel
}

func (a *Agent) Dispatch(msg message.Message) error {
	return nil
}

type AgentSession struct {
	ID      string
	session socks.Session
}

type CommandChannel struct {
	resultChan chan message.Result
	commander  Commander
}

func ()
