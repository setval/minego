package core

import (
	"fmt"
	"github.com/DiscoreMe/minego/protocol/packet"
	"github.com/DiscoreMe/minego/server"
)

func HandlerHandshake(c *server.Client) error {
	var p packet.Handshake
	if err := c.Decode(&p); err != nil {
		return err
	}
	fmt.Printf("Protocol Version: %d\nServer Address: %s\nServer Port: %d\nNext state: %d\n\n", p.ProtoVersion, p.ServerAddress, p.ServerPort, p.NextState)
	return nil
}
