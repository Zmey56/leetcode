package toHex

//Given an integer num, return a string representing its hexadecimal representation. For negative integers,
//two’s complement method is used.
//
//All the letters in the answer string should be lowercase characters,
//and there should not be any leading zeros in the answer except for the zero itself.
//
//Note: You are not allowed to use any built-in library method to directly solve this problem.

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var hex = "0123456789abcdef"
	var res string
	for num != 0 && len(res) < 8 {
		res = string(hex[num&0xf]) + res
		num >>= 4
	}
	return res

}
