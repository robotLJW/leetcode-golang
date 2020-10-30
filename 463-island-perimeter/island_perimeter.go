package _63_island_perimeter

func islandPerimeter(grid [][]int) int {
	row, column := len(grid), len(grid[0])
	var sum int
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 1 {
				sum = sum + calculate(i, j, row, column, grid)
			}
		}
	}
	return sum
}

func calculate(currentR, currentC, row, column int, grid [][]int) int {
	var ans int
	if currentR < 0 || currentR >= row || currentC < 0 || currentC >= column {
		return ans
	}
	ans = 4
	if currentR-1 >= 0 && currentR-1 < row && grid[currentR-1][currentC] == 1 {
		ans--
	}
	if currentR+1 >= 0 && currentR+1 < row && grid[currentR+1][currentC] == 1 {
		ans--
	}
	if currentC-1 >= 0 && currentC-1 < column && grid[currentR][currentC-1] == 1 {
		ans--
	}
	if currentC+1 >= 0 && currentC+1 < column && grid[currentR][currentC+1] == 1 {
		ans--
	}
	return ans
}
