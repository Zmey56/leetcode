package constructRectangle

import "math"

// A web developer needs to know how to design a web page's size. So, given a specific rectangular web page’s area,
// your job by now is to design a rectangular web page, whose length L and width W satisfy the following requirements:
//
// The area of the rectangular web page you designed must equal to the given target area.
// The width W should not be larger than the length L, which means L >= W.
// The difference between length L and width W should be as small as possible.
// Return an array [L, W] where L and W are the length and width of the web page you designed in sequence.
func constructRectangle(area int) []int {
	num := math.Sqrt(float64(area))

	var l, w = int(num), int(num)

	for l <= area {
		for j := w; j >= 1; j-- {
			if j*l == area {
				return []int{l, j}
			} else if j*l < area {
				break
			}
		}
		l++
		if l*w > area {
			w--
		}
	}
	return []int{}
}
