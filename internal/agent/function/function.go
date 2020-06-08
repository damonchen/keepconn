package function

import (
	"errors"
	"sync"

	"github.com/damonchen/keepconn/internal/message"
)

// Function function
type Function interface {
	Invoke(msg message.Message, resultChan chan message.Result) error
}

var (
	mutex    sync.Mutex
	funcMaps map[string]Function
)

var (
	ErrAlreadyExists = errors.New("already exists")
)

func Registry(typ string, fun Function) error {
	mutex.Lock()
	defer mutex.Unlock()

	if ok, _ := funcMaps[typ]; ok {
		return ErrAlreadyExists
	}

	funcMaps[typ] = fun
	return nil
}

func GetFunction(typ string) Function {
	mutex.Lock()
	defer mutex.Unlock()

	return funcMaps[typ]
}
