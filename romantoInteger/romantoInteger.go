package main

import "log"

func main() {
	log.Println(romanToInt("III"))
	log.Println(romanToInt("LVIII"))
	log.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	// Создаем словарь значений символов
	values := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	// Инициализируем переменную для хранения результата
	result := 0

	// Идем по строке справа налево
	for i := len(s) - 1; i >= 0; i-- {
		// Получаем значение текущего символа
		val := values[s[i]]

		// Проверяем, можно ли вычесть предыдущий символ из текущего
		if i > 0 && values[s[i-1]] < val {
			result += val - values[s[i-1]]
			i-- // пропускаем предыдущий символ, так как мы уже его использовали
		} else {
			result += val
		}
	}

	// Возвращаем результат
	return result
}
