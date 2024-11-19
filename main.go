package main

import (
	"bufio"
	"fmt"
	"github.com/fajarachmadyusup13/NumberOfLakes/utils"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	width, err := utils.GetUserInputInt(reader)
	if err != nil {
		logrus.Error(err)
		return
	}

	height, err := utils.GetUserInputInt(reader)
	if err != nil {
		logrus.Error(err)
		return
	}

	grid, err := utils.GetUserInputGrid(reader, height, width)
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Printf("TOTAL: %d", numOfLakes(grid))
}

// numOfLakes counter for lakes number
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

// neighbourCheck for checking neighbour of each node, and it's done recursively
func neighbourCheck(i, j int, grid [][]string) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != "." {
		return
	}
	// if a node was visited, remark it as land or change it from '.' to '#', so it would be not visited again on future iteration
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
