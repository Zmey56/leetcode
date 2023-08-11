package main

import (
	"fmt"
	"strings"
)

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))

	digits = ""
	fmt.Println(letterCombinations(digits))

	digits = "2"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	letterPhone := make(map[string][]string)
	letterPhone["1"] = []string{}
	letterPhone["2"] = []string{"a", "b", "c"}
	letterPhone["3"] = []string{"d", "e", "f"}
	letterPhone["4"] = []string{"g", "h", "i"}
	letterPhone["5"] = []string{"j", "k", "l"}
	letterPhone["6"] = []string{"m", "n", "o"}
	letterPhone["7"] = []string{"p", "q", "r", "s"}
	letterPhone["8"] = []string{"t", "u", "v"}
	letterPhone["9"] = []string{"w", "x", "y", "z"}

	result := []string{}

	if len(digits) == 1 {
		return letterPhone[digits]
	} else if len(digits) == 2 {
		for _, j := range letterPhone[string(digits[0])] {
			for _, l := range letterPhone[string(digits[1])] {
				tmp := strings.Join([]string{j, l}, "")
				result = append(result, tmp)
			}
		}
	} else if len(digits) == 3 {
		for _, j := range letterPhone[string(digits[0])] {
			for _, l := range letterPhone[string(digits[1])] {
				for _, k := range letterPhone[string(digits[2])] {
					tmp := strings.Join([]string{j, l, k}, "")
					result = append(result, tmp)
				}
			}
		}
	} else if len(digits) == 4 {
		for _, j := range letterPhone[string(digits[0])] {
			for _, l := range letterPhone[string(digits[1])] {
				for _, k := range letterPhone[string(digits[2])] {
					for _, h := range letterPhone[string(digits[3])] {
						tmp := strings.Join([]string{j, l, k, h}, "")
						result = append(result, tmp)
					}
				}
			}
		}
	}

	return result
}
