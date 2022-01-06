package net

import (
	"net"

	"github.com/bluebedmc/proto"
)

// Conn is a Minecraft Conn.
type Conn struct {
	net.Conn
	threshold int
}

// WrapConn warps an net.Conn to a Minecraft Conn.
func WrapConn(conn net.Conn) *Conn {
	return &Conn{
		Conn:      conn,
		threshold: -1,
	}
}

// ReadPacket reads a RawPacket from Conn.
func (c *Conn) ReadRawPacket(packet *proto.RawPacket) error {
	return packet.Unpack(c, c.threshold)
}

// WritePacket writes a RawPacket to Conn.
func (c *Conn) WriteRawPacket(packet proto.RawPacket) error {
	return packet.Pack(c, c.threshold)
}

// WritePacket writes a Packet to Conn.
func (c *Conn) WritePacket(packet proto.Packet) error {
	p := proto.NewRawPacket()
	err := packet.ToRaw(p)
	if err != nil {
		return err
	}

	return c.WriteRawPacket(*p)
}

// SetThreshold set threshold to Conn.
// The data packet with length equal or longer then threshold.
// will be compressed when sending.
func (c *Conn) SetThreshold(threshold int) {
	c.threshold = threshold
}
