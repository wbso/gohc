package str

func IsAllCharacterAscii(chars string) bool {
	for _, char := range chars {
		if !IsAscii(char) {
			return false
		}
	}
	return true
}

func IsAscii(char rune) bool {
	if char >= 0x20 && char <= 0x7e {
		return true
	}

	return false
}
