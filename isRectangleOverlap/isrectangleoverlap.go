package isRectangleOverlap

//An axis-aligned rectangle is represented as a list [x1, y1, x2, y2], where (x1, y1) is the coordinate of its bottom-left corner,
//and (x2, y2) is the coordinate of its top-right corner. Its top and bottom edges are parallel to the X-axis, and its left and right edges are parallel to the Y-axis.
//
//Two rectangles overlap if the area of their intersection is positive. To be clear, two rectangles that only touch at the corner or edges do not overlap.
//
//Given two axis-aligned rectangles rec1 and rec2, return true if they overlap, otherwise return false.

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if rec1[2] > rec2[0] && rec1[3] > rec2[1] && rec1[0] < rec2[2] && rec1[1] < rec2[3] {
		return true
	}
	return false
}
