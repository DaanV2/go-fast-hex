package gofasthex_test

import (
	"testing"

	gofasthex "github.com/DaanV2/go-fast-hex"
	"github.com/stretchr/testify/require"
)

func Test_RandomStrings(t *testing.T) {
	strs := []string{
		"Hello World",
		"Hello World!",
		"-/.,<>?;':\"[]{}\\|`~!@#$%^&*()_+=1234567890",
		"A quick brown fox jumps over the lazy dog",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	}

	for _, str := range strs {
		t.Run(str, func(t *testing.T) {
			higher := gofasthex.ToHigherHex([]byte(str))
			lower := gofasthex.ToLowerHex([]byte(str))

			require.Equal(t, len(str)*2, len(higher))
			require.Equal(t, len(str)*2, len(lower))

			require.Equal(t, lower, gofasthex.ToHex([]byte(str), false))
			require.Equal(t, higher, gofasthex.ToHex([]byte(str), true))

			bytes, err := gofasthex.FromHex(higher)
			require.NoError(t, err)

			require.Equal(t, str, string(bytes))
		})
	}
}

func Test_Empty(t *testing.T) {
	require.Equal(t, "", gofasthex.ToHex([]byte{}, false))
	require.Equal(t, "", gofasthex.ToHex([]byte{}, true))
	require.Equal(t, "", gofasthex.ToHigherHex([]byte{}))
	require.Equal(t, "", gofasthex.ToLowerHex([]byte{}))

	bytes, err := gofasthex.FromHex("")
	require.NoError(t, err)

	require.Equal(t, "", string(bytes))
}

func Test_Nil(t *testing.T) {
	require.Equal(t, "", gofasthex.ToHex(nil, false))
	require.Equal(t, "", gofasthex.ToHex(nil, true))
	require.Equal(t, "", gofasthex.ToHigherHex(nil))
	require.Equal(t, "", gofasthex.ToLowerHex(nil))

	bytes, err := gofasthex.FromHex("")
	require.NoError(t, err)

	require.Equal(t, "", string(bytes))
}

func Test_Invalid_HexStrings(t *testing.T) {
	invalid := []string{
		"Hello World!",
		"48656c6c6f20576f726c6",   // Missing 1 character
		"48656c6c6f20576f726c6g1", // Invalid character
	}

	for _, str := range invalid {
		t.Run(str, func(t *testing.T) {
			_, err := gofasthex.FromHex(str)
			require.Error(t, err)

			str := err.Error()
			require.Contains(t, str, "invalid")
		})
	}
}
