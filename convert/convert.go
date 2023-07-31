package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "PAYPALISHIRING"
	//PAHNAPLSIIGYIR
	//PAHNALIGYIRPSI
	//PINALSIGYAHRPI

	num := 5
	fmt.Println(s)
	fmt.Println(convert(s, num))
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]strings.Builder, numRows)
	curRow := 0
	direction := -1

	for _, c := range s {
		rows[curRow].WriteRune(c)

		if curRow == 0 || curRow == numRows-1 {
			direction = -direction
		}

		curRow += direction
	}

	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row.String())
	}

	return result.String()

}
