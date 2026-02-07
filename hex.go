package fasthex

import "github.com/DaanV2/go-fast-hex/pkg/hexvec"

func EncodeUpperWithVec(src []byte) []byte {
	return hexvec.EncodeToUpper(src)
}

func EncodeLowerWithVec(src []byte) []byte {
	return hexvec.EncodeToLower(src)
}

func EncodeWithVec(src []byte, uppercase bool) []byte {
	return hexvec.Encode(src, uppercase)
}