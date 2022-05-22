package array

// https://leetcode.cn/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	var cellM []map[byte]struct{}
	for i := 0; i < 9; i++ {
		cellM = append(cellM, make(map[byte]struct{}))
	}
	for i := 0; i < 9; i++ {
		rowM := make(map[byte]struct{})
		colM := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			rowV := board[j][i]
			colV := board[i][j]
			k := i/3*3 + j/3

			if rowV != '.' {
				if _, ok := rowM[rowV]; ok {
					return false
				} else {
					rowM[rowV] = struct{}{}
				}

				if _, ok := cellM[k][rowV]; ok {
					return false
				} else {
					cellM[k][rowV] = struct{}{}
				}
			}

			if colV != '.' {
				if _, ok := colM[colV]; ok {
					return false
				} else {
					colM[colV] = struct{}{}
				}
			}

		}
	}

	return true
}
