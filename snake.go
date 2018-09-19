package stringstyles

import (
	"strings"
)

func snakeCase(s string, sep byte) string {
	data := make([]byte, 0, len(s)*2)
	isUnderscore := true
	abbrvLen := 0

	bs := []byte(s)
	dp := 0

	for _, d := range bs {
		switch {
		case isUpper(d):
			if !isUnderscore && abbrvLen == 0 {
				data = append(data, sep)
				dp += 1
			}
			data = append(data, upperToLower(d))
			dp += 1

			abbrvLen += 1
			isUnderscore = false
		case isLower(d):
			if abbrvLen > 1 {
				data = append(data, data[dp-1])
				data[dp-1] = sep
				dp += 1
			}

			data = append(data, d)
			dp += 1

			isUnderscore = false
			abbrvLen = 0
		default:
			if !isUnderscore {
				data = append(data, sep)
				dp += 1
				isUnderscore = true
			}
			abbrvLen = 0
		}
	}

	return string(data[:dp])
}

// SnakeCase
// Punctuation is removed and spaces are replaced by single underscores
// "the_quick_brown_fox_jumps_over_the_lazy_dog"
func SnakeCase(s string) string {
	return snakeCase(s, '_')
}

// ScreamingSnakeCase
// SnakeCase with call charactor capitalised
// "SCREAMING_SNAKE_CASE"
func ScreamingSnakeCase(s string) string {
	ss := SnakeCase(s)
	return strings.ToUpper(ss)
}
