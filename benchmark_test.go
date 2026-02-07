package fasthex_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	fasthex "github.com/DaanV2/go-fast-hex"
)

var lengths = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384}

func BenchmarkVectorEncodeUpper(b *testing.B) {
	for _, length := range lengths {
		title := fmt.Sprintf("Size(%d)", length)
		src := generateData(length)

		b.Run(title, func(b *testing.B) {
			for b.Loop() {
				result := fasthex.EncodeUpper(src)
				if len(result) != length*2 {
					b.Fatalf("unexpected result length: got %d, want %d", len(result), length*2)
				}
			}
		})
	}
}

func BenchmarkVectorEncodeLower(b *testing.B) {
	for _, length := range lengths {
		title := fmt.Sprintf("Size(%d)", length)
		src := generateData(length)

		b.Run(title, func(b *testing.B) {
			for b.Loop() {
				result := fasthex.EncodeLower(src)
				if len(result) != length*2 {
					b.Fatalf("unexpected result length: got %d, want %d", len(result), length*2)
				}
			}
		})
	}
}

func BenchmarkStdEncode(b *testing.B) {
	for _, length := range lengths {
		title := fmt.Sprintf("Size(%d)", length)
		src := generateData(length)

		b.Run(title, func(b *testing.B) {
			for b.Loop() {
				result := make([]byte, hex.EncodedLen(length))
				_ = hex.Encode(result, src)
				if len(result) != length*2 {
					b.Fatalf("unexpected result length: got %d, want %d", len(result), length*2)
				}
			}
		})
	}
}

func generateData(length int) []byte {
	src := make([]byte, length)
	for i := range length {
		src[i] = byte(i % 256)
	}
	return src
}
