package reverseOnlyLetters

// Given a string s, reverse the string according to the following rules:
//
//All the characters that are not English letters remain in the same position.
//All the English letters (lowercase or uppercase) should be reversed.
//Return s after reversing it.

func reverseOnlyLetters(s string) string {
	b := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		for !((65 <= b[i] && 90 >= b[i]) || (97 <= b[i] && 122 >= b[i])) {
			i++
			if i >= j {
				return string(b)
			}
		}
		for !((65 <= b[j] && 90 >= b[j]) || (97 <= b[j] && 122 >= b[j])) {
			j--
			if i >= j {
				return string(b)
			}
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
