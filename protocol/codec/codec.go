package codec

import (
	"errors"
	"io"
)

var (
	ErrCodecVarIntTooBig   = errors.New("varInt is too big")
	ErrCodecVarLongTooBig  = errors.New("varLong is too big")
	ErrCodecStringTooSmall = errors.New("string is too small")
	ErrCodecStringTooBig   = errors.New("string is too big")
)

type Codec interface {
	Encode(io.Writer) error
	Decode(io.Reader) error
}
