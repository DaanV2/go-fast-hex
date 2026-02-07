package tabled

func EncodeUpper(src []byte) []byte {
	result := make([]byte, len(src)*2)
	k := 0
	for _, b := range src {
		v := kEncodingUpper[b]
		result[k] = v[0]
		result[k+1] = v[1]
		k += 2
	}

	return result
}

func EncodeLower(src []byte) []byte {
	result := make([]byte, len(src)*2)
	k := 0
	for _, b := range src {
		v := kEncodingLower[b]
		result[k] = v[0]
		result[k+1] = v[1]
		k += 2
	}

	return result
}