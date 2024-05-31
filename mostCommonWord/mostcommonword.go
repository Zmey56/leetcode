package mostCommonWord

import "strings"

//Given a string paragraph and a string array of the banned words banned, return the most frequent word that is not banned.
//It is guaranteed there is at least one word that is not banned, and that the answer is unique.
//
//The words in paragraph are case-insensitive and the answer should be returned in lowercase.

func mostCommonWord(paragraph string, banned []string) string {
	s := strings.ToLower(paragraph) + " "

	str := make(map[string]int)
	s1 := ""
	for _, v := range s {
		if v == ' ' || strings.Contains("!?',;.", string(v)) {
			if len(s1) != 0 {
				str[s1]++
				s1 = ""
			}
		} else {
			s1 += string(v)
		}
	}

	for _, v := range banned {
		delete(str, v)
	}

	res := ""
	max := 0
	for k, v := range str {
		if v > max {
			max = v
			res = k
		}
	}

	return res
}
