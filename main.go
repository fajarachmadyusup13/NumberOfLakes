package main

import "fmt"

func main() {
	grid := [][]string{
		{"#", ".", "#", "#", "#"},
		{".", ".", "#", "#", "#"},
		{"#", "#", ".", "#", "#"},
		{"#", "#", "#", "#", "."},
	}
	fmt.Printf("TOTAL: %d", numOfLakes(grid))
}

func numOfLakes(grid [][]string) int {

	totalLakes := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "." {
				totalLakes++
				neighbourCheck(i, j, grid)
			}
		}
	}

	return totalLakes
}

func neighbourCheck(i, j int, grid [][]string) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != "." {
		return
	}
	grid[i][j] = "#"
	neighbourCheck(i+1, j, grid)   // bottom neighbour
	neighbourCheck(i-1, j, grid)   // top neighbour
	neighbourCheck(i, j+1, grid)   // right neighbour
	neighbourCheck(i, j-1, grid)   // left neighbour
	neighbourCheck(i+1, j+1, grid) // bottom right neighbour
	neighbourCheck(i+1, j-1, grid) // bottom left neighbour
	neighbourCheck(i-1, j+1, grid) // top right neighbour
	neighbourCheck(i-1, j-1, grid) // top left neighbour
}
