package main

import (
	"strconv"
)

func compress(chars []byte) int {
	writeIndex := 0 // Index to write compressed characters
	readIndex := 0  // Index to read original characters

	for readIndex < len(chars) {
		char := chars[readIndex]
		count := 0

		// Count consecutive repeating characters
		for readIndex < len(chars) && chars[readIndex] == char {
			readIndex++
			count++
		}

		// Write compressed character
		chars[writeIndex] = char
		writeIndex++

		// Write compressed count if count > 1
		if count > 1 {
			countStr := strconv.Itoa(count)
			for _, digit := range countStr {
				chars[writeIndex] = byte(digit)
				writeIndex++
			}
		}
	}

	return writeIndex
}
