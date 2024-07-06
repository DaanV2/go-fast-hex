package gofasthex

import (
	"errors"

	hex_bytes "github.com/DaanV2/go-fast-hex/internal/bytes"
)

var (
	ErrInvalidLength = errors.New("invalid length")
)

// ToHex converts a byte slice to a hex string
func ToHex(values []byte, uppercase bool) string {
	if uppercase {
		return ToHigherHex(values)
	} else {
		return ToLowerHex(values)
	}
}

// ToHigherHex converts a byte slice to a hex string with uppercase characters: 0-9, A-F
func ToHigherHex(value []byte) string {
	l := len(value)
	if l == 0 {
		return ""
	}

	result := make([]rune, l*2)
	for i, v := range value {
		j := i * 2
		result[j] = hex_bytes.ToHigherHex(v >> 4)
		result[j+1] = hex_bytes.ToHigherHex(v)
	}

	return string(result)
}

// ToLowerHex converts a byte slice to a hex string with lowercase characters: 0-9, a-f
func ToLowerHex(value []byte) string {
	if len(value) == 0 {
		return ""
	}

	result := make([]rune, len(value)*2)
	for i, v := range value {
		j := i * 2
		result[j] = hex_bytes.ToLowerHex(v >> 4)
		result[j+1] = hex_bytes.ToLowerHex(v)
	}

	return string(result)
}

// FromHex converts a hex string to a byte slice
func FromHex(value string) ([]byte, error) {
	if err := IsValid(value); err != nil {
		return nil, err
	}

	return FromHexUnsafe(value), nil
}

// FromHex converts a hex string to a byte slice, assumes it valid length of 2 characters per byte and valid characters
// Example:
// 	FromHexUnsafe("48656c6c6f20576f726c64") -> []byte("Hello World")
func FromHexUnsafe(value string) []byte {
	l := len(value)

	runes := []rune(value)
	result := make([]byte, l/2)
	i := 0
	j := 0

	for i < l {
		a := hex_bytes.FromHex(runes[i])
		i++
		b := hex_bytes.FromHex(runes[i])
		i++
		result[j] = (a << 4) | b
		j++
	}

	return result
}

// IsValid checks if a string is a valid hex string
func IsValid(value string) error {
	if len(value)%2 != 0 {
		return ErrInvalidLength
	}

	var err error
	for _, r := range value {
		if !hex_bytes.IsValid(r) {
			err = errors.Join(err, &InvalidHexCharacterError{Character: r})
		}
	}
	return err
}
