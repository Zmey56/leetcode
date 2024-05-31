package floodFill

//An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.
//
//You are also given three integers sr, sc, and color. You should perform a flood fill on the image starting from the pixel image[sr][sc].
//
//To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of
//the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color), and so on.
//Replace the color of all of the aforementioned pixels with color.
//
//Return the modified image after performing the flood fill.

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	floodFillHelper(image, sr, sc, image[sr][sc], color)
	return image

}

func floodFillHelper(image [][]int, sr int, sc int, oldColor int, newColor int) {
	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) || image[sr][sc] != oldColor {
		return
	}
	image[sr][sc] = newColor
	floodFillHelper(image, sr+1, sc, oldColor, newColor)
	floodFillHelper(image, sr-1, sc, oldColor, newColor)
	floodFillHelper(image, sr, sc+1, oldColor, newColor)
	floodFillHelper(image, sr, sc-1, oldColor, newColor)
}
