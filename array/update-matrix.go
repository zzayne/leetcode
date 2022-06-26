package array

func updateMatrix(mat [][]int) [][]int {
	rLen, cLen := len(mat), len(mat[0])
	var queue [][]int

	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		for _, v := range direction {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < rLen && y >= 0 && y < cLen && mat[x][y] == -1 {
				mat[x][y] = mat[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}
	return mat
}
