package hex_bytes_test

import (
	"fmt"
	"slices"
	"testing"

	hex_bytes "github.com/DaanV2/go-fast-hex/internal/bytes"
	"github.com/stretchr/testify/require"
)

func Test_ToLowerHex(t *testing.T) {
	require.Equal(t, '0', hex_bytes.ToLowerHex(0))
	require.Equal(t, '1', hex_bytes.ToLowerHex(1))
	require.Equal(t, '2', hex_bytes.ToLowerHex(2))
	require.Equal(t, '3', hex_bytes.ToLowerHex(3))
	require.Equal(t, '4', hex_bytes.ToLowerHex(4))
	require.Equal(t, '5', hex_bytes.ToLowerHex(5))
	require.Equal(t, '6', hex_bytes.ToLowerHex(6))
	require.Equal(t, '7', hex_bytes.ToLowerHex(7))
	require.Equal(t, '8', hex_bytes.ToLowerHex(8))
	require.Equal(t, '9', hex_bytes.ToLowerHex(9))
	require.Equal(t, 'a', hex_bytes.ToLowerHex(10))
	require.Equal(t, 'b', hex_bytes.ToLowerHex(11))
	require.Equal(t, 'c', hex_bytes.ToLowerHex(12))
	require.Equal(t, 'd', hex_bytes.ToLowerHex(13))
	require.Equal(t, 'e', hex_bytes.ToLowerHex(14))
	require.Equal(t, 'f', hex_bytes.ToLowerHex(15))
}

func Test_ToLowerHex_Range(t *testing.T) {
	for i := range 256 {
		t.Run(fmt.Sprintf("ToLowerHex(%d)", i), func(t *testing.T) {
			value := byte(i)
			char := hex_bytes.ToLowerHex(value)

			require.True(t, hex_bytes.IsValid(char), string(char))

			conv := hex_bytes.FromHex(char)

			require.Equal(t, conv, value&0b00001111)
		})
	}
}

func Test_ToHigherHex(t *testing.T) {
	require.Equal(t, '0', hex_bytes.ToHigherHex(0))
	require.Equal(t, '1', hex_bytes.ToHigherHex(1))
	require.Equal(t, '2', hex_bytes.ToHigherHex(2))
	require.Equal(t, '3', hex_bytes.ToHigherHex(3))
	require.Equal(t, '4', hex_bytes.ToHigherHex(4))
	require.Equal(t, '5', hex_bytes.ToHigherHex(5))
	require.Equal(t, '6', hex_bytes.ToHigherHex(6))
	require.Equal(t, '7', hex_bytes.ToHigherHex(7))
	require.Equal(t, '8', hex_bytes.ToHigherHex(8))
	require.Equal(t, '9', hex_bytes.ToHigherHex(9))
	require.Equal(t, 'A', hex_bytes.ToHigherHex(10))
	require.Equal(t, 'B', hex_bytes.ToHigherHex(11))
	require.Equal(t, 'C', hex_bytes.ToHigherHex(12))
	require.Equal(t, 'D', hex_bytes.ToHigherHex(13))
	require.Equal(t, 'E', hex_bytes.ToHigherHex(14))
	require.Equal(t, 'F', hex_bytes.ToHigherHex(15))
}

func Test_ToHigherHex_Range(t *testing.T) {
	for i := range 256 {
		t.Run(fmt.Sprintf("ToHigherHex(%d)", i), func(t *testing.T) {
			value := byte(i)
			char := hex_bytes.ToHigherHex(value)

			require.True(t, hex_bytes.IsValid(char), string(char))

			conv := hex_bytes.FromHex(char)

			require.Equal(t, conv, value&0b00001111)
		})
	}
}

func Test_IsValid(t *testing.T) {
	valid := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F',
		'a', 'b', 'c', 'd', 'e', 'f',
	}

	for r := range rune(255) {
		t.Run(fmt.Sprintf("IsValid(%d)", r), func(t *testing.T) {
			expected := slices.Contains(valid, r)

			require.Equal(t, hex_bytes.IsValid(r), expected)
		})
	}
}
