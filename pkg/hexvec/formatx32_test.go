package hexvec_test

import (
	"fmt"
	"simd/archsimd"

	"github.com/DaanV2/go-fast-hex/pkg/hexvec"
)

func ExampleToLowerCharacterCodex32() {
	data := make([]uint8, 64)

	for i := range data {
		data[i] = uint8(i) & 0x0F // Keep only the last 4 bits to ensure values are in the range 0-15
	}

	v := archsimd.LoadUint8x32Slice(data)

	result := hexvec.ToLowerCharacterCodex32(v)

	var str [32]uint8
	result.Store(&str)

	fmt.Println("result", string(str[:]))
	// Output: result 0123456789abcdef0123456789abcdef
}
