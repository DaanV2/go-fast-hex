package hexvec

import "simd/archsimd"

func CanEncodeX16(l int) bool {
	return l >= Lenx16 && X16Enabled()
}

func X16Enabled() bool {
	return archsimd.X86.AVX()
}

func EncodeToUpperHexx16(src []byte) []byte {
	hexdata := make([]byte, len(src)*2)
	step := Lenx16 / 2
	var i int

	for i = 0; i < len(src) - step; i += step {
		encodeUpperx16(src[i:], hexdata[i*2:])
	}

	// Process the last remaining bytes if any
	if i != len(src) {
		i = len(src) - step
		encodeUpperx16(src[i:], hexdata[i*2:])
	}

	return hexdata
}

func EncodeToLowerHexx16(src []byte) []byte {
	hexdata := make([]byte, len(src)*2)
	step := Lenx16 / 2
	var i int

	for i = 0; i < len(src) - step; i += step {
		encodeLowerx16(src[i:], hexdata[i*2:])
	}

	// Process the last remaining bytes if any
	if i != len(src) {
		i = len(src) - step
		encodeLowerx16(src[i:], hexdata[i*2:])
	}

	return hexdata
}

func encodeUpperx16(src []byte, dst []byte) {
	step := Lenx16 / 2
	var buf [Lenx16]uint8
	k := 0

	for _, v := range src[:step] {
		buf[k] = v >> 4
		buf[k+1] = v & 0x0F
		k += 2
	}
	v := archsimd.LoadUint8x16(&buf)
	v = ToLowerCharacterCodex16(v)

	v.Store(&buf)
	copy(dst, buf[:])
}

func encodeLowerx16(src []byte, dst []byte) {
	step := Lenx16 / 2
	var buf [Lenx16]uint8
	k := 0

	for _, v := range src[:step] {
		buf[k] = v >> 4
		buf[k+1] = v & 0x0F
		k += 2
	}
	v := archsimd.LoadUint8x16(&buf)
	v = ToLowerCharacterCodex16(v)

	v.Store(&buf)
	copy(dst, buf[:])
}
