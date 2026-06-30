package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := generateRandomNumber()

	printSum((addNumbers(a, b)))

}

func generateRandomNumber() (int, int) {
	return rand.Intn((10)), rand.Intn((10))
}

func addNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func printSum(num int) {
	fmt.Printf("this is the sum of numbers : %v", num)
}
