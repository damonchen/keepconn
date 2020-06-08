package socks

import "net"

// Conn connection of the net
type Conn struct {
	net.Conn
}

// NewConn new connection
func NewConn() *Conn {
	return nil
}
