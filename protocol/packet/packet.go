package packet

import (
	"fmt"
	"github.com/DiscoreMe/minego/protocol/codec"
	"io"
)

type Packet interface {
	ID() codec.VarInt
	Decode(io.Reader) error
}

func FindPacketByID(v codec.VarInt) Packet {
	var p Packet
	switch v {
	case 0x00:
		p = &Handshake{}
	case 0x01:
		p = &Serverbound{}
	case 0xFE:
		p = &LegacyHandshaking{}
	case 0x7A:
		p = &LegacyHandshaking{}
	default:
		fmt.Println("default")
		p = &Unknown{}
	}
	return p
}
