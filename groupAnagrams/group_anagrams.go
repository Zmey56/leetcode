package groupAnagrams

// 49. Group Anagrams

//Given an array of strings strs, group the anagrams together. You can return the answer in any order.

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	groups := make(map[[26]byte][]string, len(strs))

	for _, s := range strs {
		var key [26]byte
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, g := range groups {
		result = append(result, g)
	}

	return result
}
