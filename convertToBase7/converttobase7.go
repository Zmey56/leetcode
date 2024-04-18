package convertToBase7

//Given an integer num, return a string of its base 7 representation.

func convertToBase7(num int) string {
	seven := ""
	if num == 0 {
		return "0"
	}
	for n := num; n != 0; n /= 7 {
		if n%7 < 0 {
			seven = string(byte((0-n)%7)+'0') + seven
			continue
		}
		if n%7 >= 0 {
			seven = string(byte(n%7)+'0') + seven
		}
	}
	if num < 0 {
		seven = "-" + seven
	}
	return seven
}
