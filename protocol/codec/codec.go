package codec

import "errors"

var (
	ErrCodecVarIntTooBig   = errors.New("varInt is too big")
	ErrCodecStringTooSmall = errors.New("string is too small")
	ErrCodecStringTooBig   = errors.New("string is too big")
)
