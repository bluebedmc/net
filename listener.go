package net

import (
	"net"
)

// Listener is a Minecraft server listener.
type Listener struct {
	net.Listener
}

// Listen listens as TCP (but only accepts Minecraft Conn).
func Listen(addr string) (*Listener, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Listener{l}, nil
}

// Accept a Minecraft Conn.
func (l Listener) Accept() (Conn, error) {
	conn, err := l.Listener.Accept()
	return Conn{
		Conn:      conn,
		threshold: -1,
	}, err
}
