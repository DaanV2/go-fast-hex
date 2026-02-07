package hexvec

import "simd/archsimd"

// IMPORTANT: there is an gap between '9' and 'A' and between 'F' and 'a'

const Lenx32 = 32

var (
	decreaseaAx32  = archsimd.BroadcastUint8x32(CHARACTER_CODE_a - CHARACTER_CODE_A)
	decreaseto9x32 = archsimd.BroadcastUint8x32(CHARACTER_CODE_A - CHARACTER_CODE_9 - 1)
	increaseax32   = archsimd.BroadcastUint8x32(CHARACTER_CODE_a - CHARACTER_CODE_9 - 1)
	increaseAx32   = archsimd.BroadcastUint8x32(CHARACTER_CODE_A - CHARACTER_CODE_9 - 1)
	vector0x32     = archsimd.BroadcastUint8x32(CHARACTER_CODE_0)
	vector9x32     = archsimd.BroadcastUint8x32(CHARACTER_CODE_9)
	vectorax32     = archsimd.BroadcastUint8x32(CHARACTER_CODE_a)
)

// ToLowerCharacterCode converts byte in a vector to its hexadecimal character code in lowercase.
// Assumes: that each byte in v is in the range 0-15.
func ToLowerCharacterCodex32(v archsimd.Uint8x32) archsimd.Uint8x32 {
	return toUppercodex32(v, vector0x32, vector9x32, increaseax32)
}

// ToUpperCharacterCode converts byte in a vector to its hexadecimal character code in uppercase.
// Assumes: that each byte in v is in the range 0-15.
func ToUpperCharacterCodex32(v archsimd.Uint8x32) archsimd.Uint8x32 {
	return toUppercodex32(v, vector0x32, vector9x32, increaseAx32)
}

// Convert byte data to its hexadecimal character code in uppercase or lowercase depending on the increase vector provided.
func toUppercodex32(v, vector0, vector9, vincrease archsimd.Uint8x32) archsimd.Uint8x32 {
	v = v.Add(vector0) // Increase 0-9 to '0'-'9'

	// Get a mask for values > 9
	mask := v.Greater(vector9)
	maskedadd := vincrease.Masked(mask)
	return v.Add(maskedadd) // Increase values > 9 to 'A'-'F' or 'a'-'f'
}

// fromCode converts a vector of hexadecimal character codes to their corresponding byte values (0-15).
// Automatically converts from A-Z, a-z to 0-9, and returns 0 for any non-hexadecimal character.
func FromCharacterCodex32(v archsimd.Uint8x32) archsimd.Uint8x32 {
	// Move characters from 'a'-'f' to 'A'-'F'
	maskLower := v.GreaterEqual(vectorax32)
	maskedDecrease := decreaseaAx32.Masked(maskLower)
	v = v.Sub(maskedDecrease)

	// Now we only have '0'-'9' and 'A'-'F'. Convert 'A'-'F' to the characters above '9'
	maskUpper := v.Greater(vector9x32)
	decreaseUpper := decreaseto9x32.Masked(maskUpper)
	v = v.Sub(decreaseUpper)

	// Now we have '0'-('9'+6), now to byte values
	return v.Sub(vector9x32) // Convert '0'-'9' to 0-9 and 'A'-'F' to 10-15
}
