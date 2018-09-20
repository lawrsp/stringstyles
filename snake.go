package stringstyles

import (
	"strings"
)

func snakeCase(s string, sep byte) string {
	data := make([]byte, 0, len(s)*2)
	lastIsUnderscore := true
	abbrvLen := 0
	lastIsNumber := false

	bs := []byte(s)
	dp := 0

	for _, d := range bs {
		switch {
		case isUpper(d):
			// ..X => .._x
			// ..XX => .._xx
			if !lastIsUnderscore && abbrvLen == 0 {
				data = append(data, sep)
				dp += 1
			}
			data = append(data, upperToLower(d))
			dp += 1

			abbrvLen += 1
			lastIsNumber = false
			lastIsUnderscore = false
		case isLower(d):
			// ..x => ..x
			// ..XYy => ..x_yy
			// ..123x => ...123_x
			if abbrvLen > 1 {
				data = append(data, data[dp-1])
				data[dp-1] = sep
				dp += 1
			}

			if lastIsNumber {
				data = append(data, sep)
				dp += 1
			}

			data = append(data, d)
			dp += 1

			abbrvLen = 0
			lastIsNumber = false
			lastIsUnderscore = false
		case isNumber(d):
			// ...1 => ..._1
			// ...123 => ..._123
			if !lastIsUnderscore && !lastIsNumber {
				data = append(data, sep)
				dp += 1
			}

			data = append(data, d)
			dp += 1

			abbrvLen = 0
			lastIsNumber = true
			lastIsUnderscore = false
		default:
			if !lastIsUnderscore {
				data = append(data, sep)
				dp += 1
				lastIsUnderscore = true
			}
			abbrvLen = 0
			lastIsNumber = false
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
