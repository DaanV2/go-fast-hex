package hexvec

import "simd/archsimd"

func CanEncodeX64(l int) bool {
	return l >= Lenx64 && X64Enabled()
}

func X64Enabled() bool {
	return archsimd.X86.AVX512()
}

func EncodeToUpperHexx64(src []byte) []byte {
	hexdata := make([]byte, len(src)*2)
	step := Lenx64 / 2
	var i int

	for i = 0; i < len(src) - step; i += step {
		encodeUpperx64(src[i:], hexdata[i*2:])
	}

	// Process the last remaining bytes if any
	if i != len(src) {
		i = len(src) - step
		encodeUpperx64(src[i:], hexdata[i*2:])
	}

	return hexdata
}

func EncodeToLowerHexx64(src []byte) []byte {
	hexdata := make([]byte, len(src)*2)
	step := Lenx64 / 2
	var i int

	for i = 0; i < len(src) - step; i += step {
		encodeLowerx64(src[i:], hexdata[i*2:])
	}

	// Process the last remaining bytes if any
	if i != len(src) {
		i = len(src) - step
		encodeLowerx64(src[i:], hexdata[i*2:])
	}

	return hexdata
}

func encodeUpperx64(src []byte, dst []byte) {
	step := Lenx64 / 2
	var buf [Lenx64]uint8
	k := 0

	for _, v := range src[:step] {
		buf[k] = v >> 4
		buf[k+1] = v & 0x0F
		k += 2
	}
	v := archsimd.LoadUint8x64(&buf)
	v = ToLowerCharacterCodex64(v)

	v.Store(&buf)
	copy(dst, buf[:])
}

func encodeLowerx64(src []byte, dst []byte) {
	step := Lenx64 / 2
	var buf [Lenx64]uint8
	k := 0

	for _, v := range src[:step] {
		buf[k] = v >> 4
		buf[k+1] = v & 0x0F
		k += 2
	}
	v := archsimd.LoadUint8x64(&buf)
	v = ToLowerCharacterCodex64(v)

	v.Store(&buf)
	copy(dst, buf[:])
}
