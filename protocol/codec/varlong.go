package codec

import (
	"encoding/binary"
	"io"
)

type VarLong int64

func (v *VarLong) Decode(r io.Reader) error {
	var num int8 = 0
	var result int64 = 0
	for {
		var b uint8
		if err := binary.Read(r, binary.BigEndian, &b); err != nil {
			return err
		}

		value := b & 0x7F
		result |= int64(uint64(value) << uint64(7*num))

		num++
		if num > 10 {
			return ErrCodecVarLongTooBig
		}

		if b&0x80 == 0 {
			break
		}
	}
	*v = VarLong(result)
	return nil
}

func (v VarLong) Encode(w io.Writer) error {
	var num = v
	for {
		b := uint16(num & 0x7F)
		num >>= 7
		if num != 0 {
			b |= 0x80
		}
		if err := binary.Write(w, binary.BigEndian, &b); err != nil {
			return err
		}
		if num == 0 {
			break
		}
	}
	return nil
}
