// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// var reader = bufio.NewReader(os.Stdin)

// func main() {
// 	fmt.Println("Enter height : ")
// 	var height, _ = reader.ReadString('\n')

// 	fmt.Println("Enter weight : ")
// 	var weight, _ = reader.ReadString('\n')

// 	height = strings.TrimSpace(height)
// 	weight = strings.TrimSpace(weight)

// 	heightFloat, heightErr := strconv.ParseFloat(height, 64)
// 	weightFloat, weightErr := strconv.ParseFloat(weight, 64)

// 	if heightErr != nil || weightErr != nil {
// 		fmt.Println("Please enter valid numbers for height and weight.")
// 		return
// 	}

// 	heightInMeters := heightFloat / 100

// 	fmt.Println("height : ", heightFloat)
// 	fmt.Println("weight : ", weightFloat)

// 	var bmi float64 = weightFloat / (heightInMeters * heightInMeters)

// 	fmt.Printf("BMI : %.2f", bmi)

// }
