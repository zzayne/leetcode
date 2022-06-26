package array

func orangesRotting(grid [][]int) int {
	cost, fresh, rLen, cLen := 0, 0, len(grid), len(grid[0])
	var queue [][]int
	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}

		}
	}
	if fresh > 0 && len(queue) == 0 {
		return -1
	}
	if len(queue) == 0 || fresh == 0 {
		return 0
	}

	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		tmp := queue
		queue = nil
		cost++
		for _, p := range tmp {
			for _, v := range direction {
				x := p[0] + v[0]
				y := p[1] + v[1]
				if x >= 0 && x < rLen && y >= 0 && y < cLen && grid[x][y] == 1 {
					fresh--
					grid[x][y] = 2
					queue = append(queue, []int{x, y})
				}
			}
		}
	}
	if fresh == 0 {
		return cost - 1
	}
	return -1
}
