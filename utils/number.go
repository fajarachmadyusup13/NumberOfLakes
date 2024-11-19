package utils

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

// GetUserInputInt for get user input int
func GetUserInputInt(reader *bufio.Reader) (int, error) {
	inputString, err := reader.ReadString('\n')
	if err != nil {
		logrus.Error("Error when trying to read string")
		return 0, err
	}

	intpuInt, err := userInputToInt(inputString)
	if err != nil {
		logrus.Error("Error when trying to convert string to int")
		return 0, err
	}

	return intpuInt, nil
}

// userInputToInt for convert string input to int
func userInputToInt(input string) (int, error) {
	inputString := strings.Replace(input, "\n", "", -1)

	inputInt, err := strconv.Atoi(inputString)
	if err != nil {
		logrus.Errorf("Error on: %s", inputString)
		return 0, err
	}

	return inputInt, nil
}
