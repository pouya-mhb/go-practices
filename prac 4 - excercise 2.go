package main

import "fmt"

func main () {

	var r float64 = 5.0
	pi := 3.14

	s := 2 * pi * r

	fmt.Printf("For radius of %v, the circle circumference is %.2f \n", r, s)
}