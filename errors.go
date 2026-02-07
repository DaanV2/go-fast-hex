package fasthex

import "fmt"

// InvalidHexCharacterError is returned when a rune is not a valid hex character
type InvalidHexCharacterError struct {
	Character rune // The invalid character
	Index     int  // The index of the invalid character
}

func (e *InvalidHexCharacterError) Error() string {
	return fmt.Sprintf("invalid hex character '%c' at index %d", e.Character, e.Index)
}
