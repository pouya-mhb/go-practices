package main

import "fmt"

func main () {
	a := 1
	b := 2

	fmt.Println (a+b)

	var c int = a + b

	fmt.Println (c)

	var d int = c * 2

	fmt.Println (d)

	var e float64 = 3.2
	var e2 float64 = e * e

	fmt.Println (e, e2)

	var f float32 = 3.2
	var f2 float32 = f * f
	fmt.Println (f, f2)
}