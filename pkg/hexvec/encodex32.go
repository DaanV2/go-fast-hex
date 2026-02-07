package hexvec

import "simd/archsimd"

func CanEncodeX32(l int) bool {
	return l >= Lenx32 && X32Enabled()
}

func X32Enabled() bool {
	return archsimd.X86.AVX()
}

func EncodeToUpperHexx32(src []byte) []byte {
	hexdata := make([]byte, len(src)*2)
	step := Lenx32 / 2
	var i int

	for i = 0; i < len(src) - step; i += step {
		encodeUpperx32(src[i:], hexdata[i*2:])
	}

	// Process the last remaining bytes if any
	if i != len(src) {
		i = len(src) - step
		encodeUpperx32(src[i:], hexdata[i*2:])
	}

	return hexdata
}

func EncodeToLowerHexx32(src []byte) []byte {
	hexdata := make([]byte, len(src)*2)
	step := Lenx32 / 2
	var i int

	for i = 0; i < len(src) - step; i += step {
		encodeLowerx32(src[i:], hexdata[i*2:])
	}

	// Process the last remaining bytes if any
	if i != len(src) {
		i = len(src) - step
		encodeLowerx32(src[i:], hexdata[i*2:])
	}

	return hexdata
}

func encodeUpperx32(src []byte, dst []byte) {
	step := Lenx32 / 2
	var buf [Lenx32]uint8
	// k := 0

	for i, v := range src[:step] {
		buf[i * 2] = v >> 4
		buf[(i * 2) + 1] = v & 0x0F
	}
	v := archsimd.LoadUint8x32(&buf)
	v = ToUpperCharacterCodex32(v)

	v.Store(&buf)
	copy(dst, buf[:])
}

func encodeLowerx32(src []byte, dst []byte) {
	step := Lenx32 / 2
	var buf [Lenx32]uint8
	k := 0

	for _, v := range src[:step] {
		buf[k] = v >> 4
		buf[k+1] = v & 0x0F
		k += 2
	}
	v := archsimd.LoadUint8x32(&buf)
	v = ToLowerCharacterCodex32(v)

	v.Store(&buf)
	copy(dst, buf[:])
}
