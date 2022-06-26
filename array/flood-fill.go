package array

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	rowLen, colLen := len(image), len(image[0])
	oldColor := image[sr][sc]

	if oldColor == newColor {
		return image
	}

	var fillColor func(sr int, sc int)
	fillColor = func(r int, c int) {
		if r < 0 || r >= rowLen || c < 0 || c >= colLen || image[r][c] != oldColor {
			return
		}
		image[r][c] = newColor
		fillColor(r-1, c)
		fillColor(r+1, c)
		fillColor(r, c-1)
		fillColor(r, c+1)
	}

	fillColor(sr, sc)
	return image

}
