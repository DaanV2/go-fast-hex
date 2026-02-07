package hexvec

import (
	"bytes"
	"encoding/hex"
)

func EncodeToUpper(src []byte) []byte {
	switch {
	case CanEncodeX64(len(src)):
		return EncodeToUpperHexx64(src)
	case CanEncodeX32(len(src)):
		return EncodeToUpperHexx32(src)
	case CanEncodeX16(len(src)):
		return EncodeToUpperHexx16(src)
	default:
	}

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	bytes.ToUpper(dst)

	return dst
}

func EncodeToLower(src []byte) []byte {
	switch {
	case CanEncodeX64(len(src)):
		return EncodeToLowerHexx64(src)
	case CanEncodeX32(len(src)):
		return EncodeToLowerHexx32(src)
	case CanEncodeX16(len(src)):
		return EncodeToLowerHexx16(src)
	default:
	}

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	return dst
}

func Encode(src []byte, uppercase bool) []byte {
	if uppercase {
		return EncodeToUpper(src)
	} else {
		return EncodeToLower(src)
	}
}
