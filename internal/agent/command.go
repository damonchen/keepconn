package agent

import (
	"github.com/damonchen/keepconn/internal/message"
)

type Commander struct {
}

func (c *Commander) Execute(msg message.Message, resultChan chan message.Result) error {
	return nil
}
