package returner

import (
	"github.com/damonchen/keepconn/internal/message"
)

type Returner interface {
	Return(message.Result) error
}
