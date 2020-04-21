package codec

import (
	"encoding/binary"
	"io"
)

type UShort uint16

func (v *UShort) Decode(r io.Reader) (err error) {
	var a uint16
	err = binary.Read(r, binary.BigEndian, &a)
	*v = UShort(a)
	return
}

func (v UShort) Encode(w io.Writer) (err error) {
	a := uint16(v)
	err = binary.Write(w, binary.BigEndian, &a)
	return
}
