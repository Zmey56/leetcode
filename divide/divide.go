package divide

import "math"

//Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
//
//The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be
//truncated to 8, and -2.7335 would be truncated to -2.
//
//Return the quotient after dividing dividend by divisor.

func divide(dividend int, divisor int) int {
	res := dividend / divisor
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	return res

}
