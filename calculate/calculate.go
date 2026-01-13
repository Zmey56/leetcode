package calculate

//Given a string s representing a valid expression, implement a basic calculator to evaluate it, and return the result of the evaluation.

// Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().

func calculate(s string) int {
	stack := []int{}
	result := 0
	num := 0
	sign := 1

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char >= '0' && char <= '9' {
			// Собираем число из цифр
			num = num*10 + int(char-'0')
		} else if char == '+' {
			result += sign * num
			num = 0
			sign = 1
		} else if char == '-' {
			result += sign * num
			num = 0
			sign = -1
		} else if char == '(' {
			// Сохраняем текущий результат и знак в стек
			stack = append(stack, result)
			stack = append(stack, sign)
			result = 0
			sign = 1
		} else if char == ')' {
			result += sign * num
			num = 0
			// Извлекаем знак и предыдущий результат
			prevSign := stack[len(stack)-1]
			prevResult := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// Обновляем результат: старое + (знак * то что в скобках)
			result = prevResult + (prevSign * result)
		}
	}

	return result + (sign * num)
}
