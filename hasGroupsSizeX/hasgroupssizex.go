package hasGroupsSizeX

//You are given an integer array deck where deck[i] represents the number written on the ith card.
//
//Partition the cards into one or more groups such that:
//
//Each group has exactly x cards where x > 1, and
//All the cards in one group have the same integer written on them.
//Return true if such partition is possible, or false otherwise.

func hasGroupsSizeX(deck []int) bool {
	count := make(map[int]int)
	for _, num := range deck {
		count[num]++
	}

	g := -1
	for _, c := range count {
		if g == -1 {
			g = c
		} else {
			g = gcd(g, c)
		}
	}

	return g >= 2

}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}
