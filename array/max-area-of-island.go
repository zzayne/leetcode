package array

func maxAreaOfIsland(grid [][]int) int {
	rowLen, colLen, max := len(grid), len(grid[0]), 0

	var dfs func(i, j int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= rowLen || c < 0 || c >= colLen || grid[r][c] != 1 {
			return 0
		}
		count := 1
		grid[r][c] = 2
		count += dfs(r-1, c)
		count += dfs(r+1, c)
		count += dfs(r, c-1)
		count += dfs(r, c+1)
		return count
	}

	maxInt := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			max = maxInt(dfs(i, j), max)
		}
	}
	return max
}
