package backspaceCompare

import (
	"strings"
	"sync"
)

//Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.
//
//Note that after backspacing an empty text, the text will continue empty.

func backspaceCompare(s string, t string) bool {
	result := []rune{}
	backspace := "#" // backspace character
	for _, c := range s {
		if c == rune(backspace[0]) {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, c)
		}
	}

	result1 := []rune{}
	for _, c := range t {
		if c == rune(backspace[0]) {
			if len(result1) > 0 {
				result1 = result1[:len(result1)-1]
			}
		} else {
			result1 = append(result1, c)
		}
	}

	return strings.Compare(string(result), string(result1)) == 0

}

func backspaceCompareVer2(s string, t string) bool {
	resultChan := make(chan []rune, 2)
	var wg sync.WaitGroup

	wg.Add(2)
	go processString(s, resultChan, &wg)
	go processString(t, resultChan, &wg)

	wg.Wait()
	close(resultChan)

	var results [][]rune
	for result := range resultChan {
		results = append(results, result)
	}

	return strings.Compare(string(results[0]), string(results[1])) == 0
}

func processString(s string, resultChan chan<- []rune, wg *sync.WaitGroup) {
	defer wg.Done()

	result := []rune{}
	backspace := "#"
	for _, c := range s {
		if c == rune(backspace[0]) {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, c)
		}
	}

	resultChan <- result
}
