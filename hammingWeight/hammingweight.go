package hammingWeight

//Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits
//it has (also known as the Hamming weight).

func hammingWeight(num uint32) int {
	weigt := 0
	for num != 0 {
		if num&1 == 1 {
			weigt++
		}
		num >>= 1
	}
	return weigt
}
