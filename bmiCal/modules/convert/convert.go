package convert

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertToFloat(input1 string, input2 string) (float64, error, float64, error) {

	inputTrimmed1 := strings.TrimSpace(input1)
	inputTrimmed2 := strings.TrimSpace(input2)

	inputFloat1, InputErr1 := strconv.ParseFloat(inputTrimmed1, 64)
	inputFloat2, InputErr2 := strconv.ParseFloat(inputTrimmed2, 64)

	if InputErr1 != nil || InputErr2 != nil {
		fmt.Println("Please enter valid numbers for height and weight.")

	}

	return inputFloat1, InputErr1, inputFloat2, InputErr2
}
