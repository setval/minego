package packet

import (
	"github.com/DiscoreMe/minego/protocol/codec"
	"io"
)

type Unknown struct{}

func (u *Unknown) ID() codec.VarInt {
	return 0x99
}

func (u *Unknown) Decode(io.Reader) error {
	return nil
}
