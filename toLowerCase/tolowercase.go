package toLowerCase

// Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.
func toLowerCase(s string) string {
	// Create a lookup table for ASCII characters.
	lookup := [256]byte{}
	for i := 0; i < 256; i++ {
		c := byte(i)
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A' // Convert to lowercase.
		}
		lookup[i] = c
	}

	// Convert the string to a byte slice and apply the lookup table.
	b := []byte(s)
	for i, c := range b {
		b[i] = lookup[c]
	}

	return string(b)
}
