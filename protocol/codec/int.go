package codec

import (
	"encoding/binary"
	"io"
)

type VarInt int32

func (v *VarInt) Decode(r io.Reader) error {
	var num int8 = 0
	var result int32 = 0
	for {
		var b uint8
		if err := binary.Read(r, binary.BigEndian, &b); err != nil {
			return err
		}

		value := b & 0x7F
		result |= int32(uint(value) << uint(7*num))

		num++
		if num > 5 {
			return ErrCodecVarIntTooBig
		}

		if b&0x80 == 0 {
			break
		}
	}
	*v = VarInt(result)
	return nil
}

func (v VarInt) Encode(w io.Writer) error {
	var num = v
	for {
		b := num & 0x7F
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
