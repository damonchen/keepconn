package message

import (
	"bytes"
	"io"
)

const (
	// MsgUnknown unknown message
	MsgUnknown = iota
	// MsgResult result message
	MsgResult
	// MsgAuth authenticate message
	MsgAuth
	// MsgData data message
	MsgData
	// MsgAck ack message
	MsgAck
)

// for simple, using json for exchange message
type Marshler interface {
	Marshal(v interface{}) (io.Reader, error)
}

type Unmarshaler interface {
	Unmarshal(io.Reader, interface{}) error
}

// Frame frame
type Frame interface {
	// GetType get message type
	Type() string
	// GetSize get size
	GetSize() uint16
	// Packed packed
	Packed() (buf *bytes.Buffer, err error)
	// Unpack unpack
	Unpack(r io.Reader) error
}

type Message Frame

type Result interface {
	Frame
}

type CommandMessage struct {
}
