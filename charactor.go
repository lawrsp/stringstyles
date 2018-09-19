package stringstyles

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func upperToLower(c byte) byte {
	return c + 32
}

func lowerToUper(c byte) byte {
	return c - 32
}
