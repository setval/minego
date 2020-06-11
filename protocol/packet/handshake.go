package packet

import (
	"github.com/DiscoreMe/minego/protocol/codec"
	"io"
)

type Handshake struct {
	ProtoVersion  codec.VarInt
	ServerAddress codec.String
	ServerPort    codec.UShort
	NextState     codec.VarInt
}

func (h *Handshake) ID() codec.VarInt {
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

type LegacyHandshaking struct {
	Payload codec.UShort
}

func (h *LegacyHandshaking) ID() codec.VarInt {
	return 0xFE
}

func (h *LegacyHandshaking) Decode(conn io.Reader) error {
	return h.Payload.Decode(conn)
}
