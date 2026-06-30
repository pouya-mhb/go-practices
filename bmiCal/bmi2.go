package main

import (
	"bmiCal/modules/bmical"
	"bmiCal/modules/convert"
	"bmiCal/modules/getInput"
)

func main() {

	height, weight := getinput.GetInput()
	heightFloat, heightErr, weightFloat, weightErr := convert.ConvertToFloat(height, weight)

	if heightErr != nil || weightErr != nil {
		return
	}

	bmical.BmiCalculator(heightFloat, weightFloat)

}
