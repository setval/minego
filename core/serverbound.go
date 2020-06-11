package core

import (
	"fmt"
	"github.com/DiscoreMe/minego/protocol/codec"
	"github.com/DiscoreMe/minego/protocol/packet"
	"github.com/DiscoreMe/minego/server"
)

func HandleServerbound(c *server.Client) error {
	var p packet.Serverbound
	if err := c.Decode(p); err != nil {
		return err
	}

	fmt.Println(1, p.Payload)
	m := codec.VarLong(67)
	return c.Write(p, &m)
}
