package main

// Function to find the greatest common divisor (GCD) of two integers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find the largest common divisor string
func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	// Calculate the greatest common divisor of lengths
	commonLength := gcd(len1, len2)

	// Get the candidate string
	candidate := str1[:commonLength]

	// Check if the candidate string can divide both str1 and str2
	if checkDivision(str1, candidate) && checkDivision(str2, candidate) {
		return candidate
	}

	return ""
}

// Function to check if a string can be formed by repeating a substring
func checkDivision(s, sub string) bool {
	subLen := len(sub)
	for i := 0; i < len(s); i += subLen {
		if i+subLen > len(s) || s[i:i+subLen] != sub {
			return false
		}
	}
	return true
}
