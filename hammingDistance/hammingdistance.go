package hammingDistance

import "log"

//The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
//
//Given two integers x and y, return the Hamming distance between them.

func hammingDistance(x int, y int) int {
	//get bits of x and y
	xBits := getBits(x)
	yBits := getBits(y)
	log.Println("xBits", xBits)
	log.Println("yBits", yBits)

	//get the length of the longest bits
	maxLen := len(xBits)
	if len(yBits) > maxLen {
		maxLen = len(yBits)
	}

	//pad the shorter bits with 0
	for len(xBits) < maxLen {
		xBits = append(xBits, 0)
	}

	for len(yBits) < maxLen {
		yBits = append(yBits, 0)

	}
	log.Println("xBits 2", xBits)
	log.Println("yBits 2", yBits)

	//compare the bits
	hammingDistance := 0
	for i := 0; i < maxLen; i++ {
		if xBits[i] != yBits[i] {
			hammingDistance++
		}

	}

	return hammingDistance

}

func getBits(x int) []int {
	bits := []int{}
	for x > 0 {
		bits = append(bits, x%2)
		x = x / 2
		log.Println("x", x, "bits", bits)
	}
	return bits
}
