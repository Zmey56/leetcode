package isOneBitCharacter

//We have two special characters:
//
//The first character can be represented by one bit 0.
//The second character can be represented by two bits (10 or 11).
//Given a binary array bits that ends with 0, return true if the last character must be a one-bit character.

func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}
	return i == len(bits)-1
}
