package main

import "strings"

func removeStars(s string) string {
	result := []string{}

	countStars := 0
	for n, i := range s {
		if string(i) != "*" {
			if countStars != 0 {
				result = result[:len(result)-countStars]
				countStars = 0
			}
			result = append(result, string(i))
		} else {
			countStars++
			if n == len(s)-1 {
				result = result[:len(result)-countStars]
			}
		}
	}
	return strings.Join(result, "")
}
