package bmical

import "fmt"

func BmiCalculator(height float64, weight float64) float64 {

	heightInMeters := height / 100

	fmt.Println("height : ", height)
	fmt.Println("weight : ", weight)

	var bmi float64 = weight / (heightInMeters * heightInMeters)

	fmt.Printf("BMI : %.2f", bmi)
	return bmi
}
