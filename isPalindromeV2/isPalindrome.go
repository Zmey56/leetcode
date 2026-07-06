package main

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Пропускаем все не-буквы и не-цифры слева
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}

		// Пропускаем все не-буквы и не-цифры справа
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// Если указатели не встретились, сравниваем символы
		if left < right {
			if toLower(s[left]) != toLower(s[right]) {
				return false
			}
			left++
			right--
		}
	}

	return true
}

// Вспомогательная функция проверки: буква ли это (A-Z, a-z) или цифра (0-9)
func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9')
}

// Вспомогательная функция приведения к нижнему регистру
func toLower(b byte) byte {
	// В ASCII заглавные буквы находятся в диапазоне 65-90 ('A'-'Z').
	// Чтобы получить строчную, достаточно прибавить 32.
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}
