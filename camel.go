package stringstyles

// CamelCase
// Spaces and punctuation are removed and the first letter of each word is capitalised.
// "theQuickBrownFoxJumpsOverTheLazyDog"
func CamelCase(s string) string {
	data := make([]byte, 0, len(s))
	bs := []byte(s)
	abbrvLen := 0
	isNoChar := false

	for i, d := range bs {
		switch {
		case isUpper(d):
			if abbrvLen == 0 && i != 0 {
				data = append(data, d)
			} else {
				data = append(data, upperToLower(d))
			}
			abbrvLen += 1
			isNoChar = false
		case isLower(d):
			if abbrvLen > 1 {
				dlen := len(data)
				data[dlen-1] = lowerToUper(data[dlen-1])
			}

			if isNoChar && i != 0 {
				data = append(data, lowerToUper(d))
			} else {
				data = append(data, d)
			}

			abbrvLen = 0
			isNoChar = false
		default:
			abbrvLen = 0
			isNoChar = true
		}
	}

	return string(data[:])
}

// PascalCase
// CamelCase with first charactor is also upper case
// "TheQuickBrownFoxJumpsOverTheLazyDog"
func PascalCase(s string) string {
	cs := CamelCase(s)
	bcs := []byte(cs)
	bcs[0] = lowerToUper(bcs[0])
	return string(bcs[:])
}
