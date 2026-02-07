package hexvec

import "simd/archsimd"

// IMPORTANT: there is an gap between '9' and 'A' and between 'F' and 'a'

const Lenx16 = 16

// ToLowerCharacterCode converts byte in a vector to its hexadecimal character code in lowercase.
// Assumes: that each byte in v is in the range 0-15.
func ToLowerCharacterCodex16(v archsimd.Uint8x16) archsimd.Uint8x16 {
	vector0 := archsimd.BroadcastUint8x16(CHARACTER_CODE_0)
	vector9 := archsimd.BroadcastUint8x16(CHARACTER_CODE_9)
	increase := archsimd.BroadcastUint8x16(CHARACTER_CODE_a - CHARACTER_CODE_9 - 1)

	return toUppercodex16(v, vector0, vector9, increase)
}

// ToUpperCharacterCode converts byte in a vector to its hexadecimal character code in uppercase.
// Assumes: that each byte in v is in the range 0-15.
func ToUpperCharacterCodex16(v archsimd.Uint8x16) archsimd.Uint8x16 {
	vector0 := archsimd.BroadcastUint8x16(CHARACTER_CODE_0)
	vector9 := archsimd.BroadcastUint8x16(CHARACTER_CODE_9)
	increase := archsimd.BroadcastUint8x16(CHARACTER_CODE_A - CHARACTER_CODE_9 - 1)

	return toUppercodex16(v, vector0, vector9, increase)
}

// ToCharacterCodex converts byte in a vector to its hexadecimal character code in uppercase or lowercase depending on the uppercase parameter.
func ToCharacterCodex16(v archsimd.Uint8x16, uppercase bool) archsimd.Uint8x16 {
	if uppercase {
		return ToUpperCharacterCodex16(v)
	} else {
		return ToLowerCharacterCodex16(v)
	}
}

// Convert byte data to its hexadecimal character code in uppercase or lowercase depending on the increase vector provided.
func toUppercodex16(v, vector0, vector9, vincrease archsimd.Uint8x16) archsimd.Uint8x16 {
	v = v.Add(vector0) // Increase 0-9 to '0'-'9'

	// Get a mask for values > 9
	mask := v.Greater(vector9)
	maskedadd := vincrease.Masked(mask)
	return v.Add(maskedadd) // Increase values > 9 to 'A'-'F' or 'a'-'f'
}

// fromCode converts a vector of hexadecimal character codes to their corresponding byte values (0-15).
// Automatically converts from A-Z, a-z to 0-9, and returns 0 for any non-hexadecimal character.
func FromCharacterCodex16(v archsimd.Uint8x16) archsimd.Uint8x16 {
	vector9 := archsimd.BroadcastUint8x16(CHARACTER_CODE_9)
	vectora := archsimd.BroadcastUint8x16(CHARACTER_CODE_a)
	decreaseaA := archsimd.BroadcastUint8x16(CHARACTER_CODE_a - CHARACTER_CODE_A)

	// Move characters from 'a'-'f' to 'A'-'F'
	maskLower := v.GreaterEqual(vectora)
	maskedDecrease := decreaseaA.Masked(maskLower)
	v = v.Sub(maskedDecrease)

	// Now we only have '0'-'9' and 'A'-'F'. Convert 'A'-'F' to the characters above '9'
	maskUpper := v.Greater(vector9)
	decreaseUpper := archsimd.BroadcastUint8x16(CHARACTER_CODE_A - CHARACTER_CODE_9 - 1).Masked(maskUpper)
	v = v.Sub(decreaseUpper)

	// Now we have '0'-('9'+6), now to byte values
	return v.Sub(vector9) // Convert '0'-'9' to 0-9 and 'A'-'F' to 10-15
}
