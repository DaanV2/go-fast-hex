package hex_bytes

const (
	c_A_RUNE = rune('A')
	c_F_RUNE = rune('F')
	c_a_RUNE = rune('a')
	c_f_RUNE = rune('f')
	c_0_RUNE = rune('0')
	c_9_RUNE = rune('9')

	c_9_A_DIFF = c_A_RUNE - c_9_RUNE
	c_9_a_DIFF = c_a_RUNE - c_9_RUNE

	LOWER_BIT_MASK = 0b00001111
)

// ToLowerHex converts the lower 4 bits of a byte to a rune: 0-9, a-f
func ToLowerHex(b byte) rune {
	c := rune(b & LOWER_BIT_MASK)
	c += c_0_RUNE

	if c > c_9_RUNE {
		c += (c_9_a_DIFF - 1)
	}

	return c
}

// ToHigherHex converts the higher 4 bits of a byte to a rune: 0-9, A-F
func ToHigherHex(b byte) rune {
	c := rune(b & LOWER_BIT_MASK)
	c += c_0_RUNE

	if c > c_9_RUNE {
		c += (c_9_A_DIFF - 1)
	}

	return c
}

// FromHex converts a rune to a byte
func FromHex(r rune) byte {
	if r >= c_a_RUNE {
		r -= (c_9_a_DIFF - 1)
	} else if r >= c_A_RUNE {
		r -= (c_9_A_DIFF - 1)
	}

	return byte(r - c_0_RUNE)
}

// IsValid checks if a rune is a valid hex character
func IsValid(r rune) bool {
	if r >= c_0_RUNE && r <= c_9_RUNE {
		return true
	}
	if r >= c_A_RUNE && r <= c_F_RUNE {
		return true
	}
	if r >= c_a_RUNE && r <= c_f_RUNE {
		return true
	}

	return false
}
