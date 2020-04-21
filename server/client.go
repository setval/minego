package server

import (
	"github.com/DiscoreMe/minego/protocol/codec"
	"github.com/DiscoreMe/minego/protocol/packet"
	"net"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) PacketID() (codec.VarInt, error) {
	var id codec.VarInt
	return id, id.Decode(c.conn)
}

func (c *Client) PackLength() (codec.VarInt, error) {
	var id codec.VarInt
	return id, id.Decode(c.conn)
}

func (c *Client) Disconnect() error {
	return c.conn.Close()
}

func (c *Client) Decode(p packet.Packet) error {
	return p.Decode(c.conn)
}
