package convertToTitle

//Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.

func convertToTitle(columnNumber int) string {
	var result string
	for columnNumber > 0 {
		columnNumber--
		remainder := columnNumber % 26
		result = string('A'+remainder) + result
		columnNumber /= 26
	}
	return result
}
