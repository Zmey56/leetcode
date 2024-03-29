package countSegments

import (
	"log"
	"strings"
)

//Given a string s, return the number of segments in the string.
//
//A segment is defined to be a contiguous sequence of non-space characters.

func countSegments(s string) int {
	count := 0
	sArr := strings.Split(s, " ")
	log.Println("S: ", s, "Sarr: ", sArr)
	for _, v := range sArr {
		log.Println("V: ", v, "String: ", v, "Count: ", count)
		v = strings.ReplaceAll(v, " ", "")
		if v != "" {
			count++
		}
	}
	return count
}
