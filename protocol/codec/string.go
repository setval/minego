package codec

import "io"

type String string

func (s *String) Decode(r io.Reader) error {
	var length VarInt
	if err := length.Decode(r); err != nil {
		return err
	}
	if length <= 0 {
		return ErrCodecStringTooSmall
	}
	if length > 131071 {
		return ErrCodecStringTooBig
	}
	var str = make([]byte, length)
	_, err := r.Read(str)
	if err != nil {
		return err
	}
	*s = String(string(str))
	return nil
}

func (s String) Encode(w io.Writer) error {
	bytes := []byte(s)
	if err := VarInt(len(bytes)).Encode(w); err != nil {
		return err
	}
	_, err := w.Write(bytes)
	return err
}
