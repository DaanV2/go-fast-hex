package tabled_test

import (
	"fmt"
	"testing"
)

func Test_GenerateEncoding(t *testing.T) {
	var kEncoding [256][2]rune

	for i := range kEncoding {
		upper := "0123456789abcdef"[i>>4]
		lower := "0123456789abcdef"[i&0x0F]
		kEncoding[i][0] = rune(upper)
		kEncoding[i][1] = rune(lower)
	}

	data := "var kEncodingLower = [256][2]byte{\n"
	for _, v := range kEncoding {
		data += fmt.Sprintf("\t{'%c', '%c'},\n", v[0], v[1])
	}
	data += "}\n"
	fmt.Println(data)
}
