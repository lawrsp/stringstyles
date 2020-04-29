package stringstyles

// KebabCase
// Similar to snake case, above, except hyphens rather than underscores are used to replace spaces.
// e.g. "the-quick-brown-fox-jumps-over-the-lazy-dog"
func KebabCase(s string) string {
	return snakeCase(s, '-')
}
