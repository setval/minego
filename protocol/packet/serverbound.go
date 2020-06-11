package packet

import (
	"github.com/DiscoreMe/minego/protocol/codec"
	"io"
)

type Serverbound struct {
	Payload codec.VarLong
}

func (s Serverbound) ID() codec.VarInt {
	return 1
}
func (s Serverbound) Decode(r io.Reader) error {
	return s.Payload.Decode(r)
}
