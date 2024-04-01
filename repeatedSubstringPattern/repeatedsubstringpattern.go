package repeatedSubstringPattern

import (
	"log"
	"strings"
)

//Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.

func repeatedSubstringPattern(s string) bool {
	t := s + s
	log.Println("LEN", t[1:len(t)-1])
	log.Println("s", s)
	return strings.Contains(t[1:len(t)-1], s)

}
