package main

import (
	"adventofcode2023/util"
	"fmt"
	"strconv"
	"strings"
)

const (
	fileName = "misc/bridgeNumbers"
	MUL      = "X"
	ADD      = "A"
)

func main() {
	lines := util.GetRowsFromFile(fileName)

	sum := 0
	for _, line := range lines {
		stringVals := strings.Split(line, ":")
		cal, _ := strconv.Atoi(stringVals[0])
		values := util.ExtractNumbersByDelimiter(stringVals[1], " ")

		fmt.Printf("Cal %d\n", cal)
		fmt.Printf("%+v\n", values)

		validLine, err := checkLine(values[0], strconv.Itoa(values[0]), cal, values[1:])
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
		} else if validLine {
			sum += cal
		}

		fmt.Printf("========================\n")
	}

	fmt.Printf("Sum of calibration: %d", sum)
}

func checkLine(currentValue int, currentFunction string, calibrationNumber int, values []int) (bool, error) {
	if len(values) == 0 {
		acceptablePath := currentValue == calibrationNumber
		if acceptablePath {
			fmt.Printf("Acceptable path: %s\n", currentFunction)
			return true, nil
		}
		fmt.Printf("invalid path: %s\n", currentFunction)
		return false, nil
	}
	multiFunc := fmt.Sprintf("%s*%d", currentFunction, values[0])
	multi := currentValue * values[0]

	validMulti, err := checkLine(multi, multiFunc, calibrationNumber, values[1:])
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else if validMulti {
		return true, nil
	}

	add := currentValue + values[0]
	addFunc := fmt.Sprintf("%s+%d", currentFunction, values[0])
	validAdd, err := checkLine(add, addFunc, calibrationNumber, values[1:])
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else if validAdd {
		return true, nil
	}

	return false, nil
}
