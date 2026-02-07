package fasthex

import "github.com/DaanV2/go-fast-hex/pkg/hexvec"

func EncodeUpper(src []byte) []byte {
	return hexvec.EncodeToUpper(src)
}

func EncodeLower(src []byte) []byte {
	return hexvec.EncodeToLower(src)
}

func Encode(src []byte, uppercase bool) []byte {
	return hexvec.Encode(src, uppercase)
}