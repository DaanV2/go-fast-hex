package hexvec_test

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/DaanV2/go-fast-hex/pkg/hexvec"
	"github.com/stretchr/testify/assert"
)

func Fuzz_EncodeLower(f *testing.F) {
	acceptableRunes := "0123456789abcdef"
	f.Add(int16(2))
	f.Add(int16(8))
	f.Add(int16(16))
	f.Add(int16(17))
	f.Add(int16(31))
	f.Add(int16(32))
	f.Add(int16(33))
	f.Add(int16(63))
	f.Add(int16(64))
	f.Add(int16(65))
	f.Add(int16(127))
	f.Add(int16(128))
	f.Add(int16(129))

	f.Fuzz(func(t *testing.T, l int16) {
		d := make([]byte, l)
		for i := range d {
			d[i] = byte(i)
		}

		encoded := hexvec.EncodeToLower(d)
		assert.Len(t, encoded, int(2*l))
		for i, v := range encoded {
			assert.True(t, strings.ContainsRune(acceptableRunes, rune(v)), "Invalid character at position %d: %c", i, v)
		}

		actual := hex.AppendEncode(nil, d)
		assert.Equal(t, string(encoded), string(actual))
	})
}
