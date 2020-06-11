package server

import (
	"bytes"
	"fmt"
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

func (c *Client) Write(p packet.Packet, data codec.Codec) error {
	var b bytes.Buffer
	if err := data.Encode(&b); err != nil {
		return err
	}

	if err := codec.VarInt(b.Len() + p.ID().Length()).Encode(c.conn); err != nil {
		return err
	}

	if err := p.ID().Encode(c.conn); err != nil {
		return err
	}

	fmt.Println(b.String())

	ss := codec.String(b.String())

	if err := ss.Encode(c.conn); err != nil {
		return err
	}
	//_, err := c.conn.Write(b.Bytes())
	//if err != nil {
	//	return err
	//}
	return nil
}

func (c *Client) Disconnect() error {
	return c.conn.Close()
}

func (c *Client) Decode(p packet.Packet) error {
	return p.Decode(c.conn)
}
