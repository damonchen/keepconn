package function

import (
	"github.com/damonchen/keepconn/internal/message"
)

type Cmd struct {
}

type CmdResponse struct {
}

func (c *Cmd) Invoke(msg message.Message, resultChan chan message.Result) error {
	cmdMsg := msg.(message.CommandMessage)

	resultChan <- nil
}
