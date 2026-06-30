package convert

import (
	"fmt"
	"strconv"
	"strings"
)

func Trimming(input string) string {
	inputTrimmed1 := strings.TrimSpace(input)
	return inputTrimmed1
}

func ConvertToFloat(input1 string, input2 string) (float64, error, float64, error) {

	inputFloat1, InputErr1 := strconv.ParseFloat(Trimming(input1), 64)
	inputFloat2, InputErr2 := strconv.ParseFloat(Trimming(input2), 64)

	if InputErr1 != nil || InputErr2 != nil {
		fmt.Println("Please enter valid numbers for height and weight.")

	}

	return inputFloat1, InputErr1, inputFloat2, InputErr2
}
