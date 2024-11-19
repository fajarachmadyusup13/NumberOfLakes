package utils

import (
	"bufio"
	"errors"
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

// GetUserInputGrid get user input grid
func GetUserInputGrid(reader *bufio.Reader, height, width int) ([][]string, error) {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		rowString, err := reader.ReadString('\n')
		if err != nil {
			logrus.Error("Error when trying to read string")
			return nil, err
		}

		cleanedRow := strings.Replace(rowString, "\n", "", -1)

		// checking is row contains invalid characters (other than # and .)
		if containsInvalidCharacter(cleanedRow) {
			logrus.Errorf("Error on: %s", cleanedRow)
			return nil, errors.New("invalid character in row")
		}

		// convert string into list
		rowArr := strings.Split(cleanedRow, "")
		if len(rowArr) != width {
			logrus.Errorf("Error on: %s", cleanedRow)
			return nil, errors.New("number of columns mismatch")
		}
		grid[i] = rowArr
	}

	return grid, nil
}

func containsInvalidCharacter(input string) bool {
	re := regexp.MustCompile(`[^#.]`)
	return re.MatchString(input)
}
