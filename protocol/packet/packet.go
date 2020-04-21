package packet

import (
	"github.com/DiscoreMe/minego/protocol/codec"
	"io"
)

type Packet interface {
	ID() int
	Decode(io.Reader) error
}

func FindPacketByID(v codec.VarInt) Packet {
	var p Packet
	switch v {
	case 0x00:
		p = &Handshake{}
	default:
		p = &Unknown{}
	}
	return p
}

type Unknown struct{}

func (u *Unknown) ID() int {
	return 0x0
}
func (u *Unknown) Decode(io.Reader) error {
	return nil
}

type Handshake struct {
	ProtoVersion  codec.VarInt
	ServerAddress codec.String
	ServerPort    codec.UShort
	NextState     codec.VarInt
}

func (h *Handshake) ID() int {
	return 0x00
}
func (h *Handshake) Decode(conn io.Reader) error {
	if err := h.ProtoVersion.Decode(conn); err != nil {
		return err
	}
	if err := h.ServerAddress.Decode(conn); err != nil {
		return err
	}
	if err := h.ServerPort.Decode(conn); err != nil {
		return err
	}
	return h.NextState.Decode(conn)
}
