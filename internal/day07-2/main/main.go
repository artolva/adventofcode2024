package main

import (
	"adventofcode2023/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	fileName = "misc/bridgeNumbers"
	MUL      = "X"
	ADD      = "A"
)

func main() {
	now := time.Now()
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

		//break
		fmt.Printf("========================\n")
	}

	fmt.Printf("Sum of calibration: %d\n", sum)
	fmt.Printf("Found in %d\n", time.Now().UnixMilli()-now.UnixMilli())
}

func checkLine(currentValue int, currentFunction string, calibrationNumber int, values []int) (bool, error) {
	if len(values) == 0 {
		acceptablePath := currentValue == calibrationNumber
		if acceptablePath {
			//fmt.Printf("Acceptable path: %s\n", currentFunction)
			return true, nil
		}
		//fmt.Printf("invalid path: %s\n", currentFunction)
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

	n := fmt.Sprintf("%d%d", currentValue, values[0])
	concat, _ := strconv.Atoi(n)
	concatFunc := fmt.Sprintf("%s%d", currentFunction, values[0])
	validConcat, err := checkLine(concat, concatFunc, calibrationNumber, values[1:])
	if err != nil {
		fmt.Println(err.Error())
	} else if validConcat {
		return true, nil
	}

	return false, nil
}

//
//func checkLine(currentValue int, currentFunction string, calibrationNumber int, values []int) (bool, error) {
//	if len(values) == 0 {
//		acceptablePath := currentValue == calibrationNumber
//		if acceptablePath {
//			fmt.Printf("Acceptable path: %s\n", currentFunction)
//			return true, nil
//		}
//		fmt.Printf("invalid path: %s\n", currentFunction)
//		return false, nil
//	}
//
//	currentConcat := ""
//	for i := 0; i < len(values); i++ {
//		currentConcat = fmt.Sprintf("%s%d", currentConcat, values[i])
//		nextValue, _ := strconv.Atoi(currentConcat)
//		multiFunc := fmt.Sprintf("%s*%d", currentFunction, nextValue)
//		multi := currentValue * nextValue
//
//		validMulti, err := checkLine(multi, multiFunc, calibrationNumber, values[i+1:])
//		if err != nil {
//			fmt.Printf("%s\n", err.Error())
//		} else if validMulti {
//			return true, nil
//		}
//
//		add := currentValue + nextValue
//		addFunc := fmt.Sprintf("%s+%d", currentFunction, nextValue)
//		validAdd, err := checkLine(add, addFunc, calibrationNumber, values[i+1:])
//		if err != nil {
//			fmt.Printf("%s\n", err.Error())
//		} else if validAdd {
//			return true, nil
//		}
//	}
//
//	return false, nil
//}
