package findContentChildren

import "sort"

//Assume you are an awesome parent and want to give your children some cookies. But, you should give each child at most one cookie.
//
//Each child i has a greed factor g[i], which is the minimum size of a cookie that the child will be content with;
//and each cookie j has a size s[j]. If s[j] >= g[i], we can assign the cookie j to the child i, and the child i will be content.
//Your goal is to maximize the number of your content children and output the maximum number.

func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(s); j++ {
			if g[i] <= s[j] {
				count++
				s = append(s[:j], s[j+1:]...)
				break
			}
		}
	}
	return count
}
