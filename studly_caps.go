package stringstyles

func getBit(src uint64, pos uint) uint64 {
	pos = pos % 64
	return src & (1 << pos)
}

// Studly caps
// Mixed case with no semantic or syntactic significance to the use of the capitals
// e.g. "tHeqUicKBrOWnFoXJUmpsoVeRThElAzydOG"
func StudlyCaps(s string, random uint64) string {
	data := make([]byte, 0, len(s))
	bs := []byte(s)

	for i, d := range bs {
		switch {
		case isUpper(d):
			if getBit(random, uint(i)) != 0 {
				d = upperToLower(d)
			}
			data = append(data, d)
		case isLower(d):
			if getBit(random, uint(i)) != 0 {
				d = lowerToUper(d)
			}
			data = append(data, d)
		default:
		}
	}

	return string(data[:])
}
