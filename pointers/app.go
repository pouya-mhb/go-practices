package main

import "fmt"

func main() {

	age := 30 // age is the variable with vale of 30
	fmt.Println(age)

	myAge := &age // myAge is the pointer of age variable
	fmt.Println(myAge)

	*myAge = 33 // change the value of the pointer of age variable
	fmt.Println(*myAge)
	fmt.Println(age) // the value of age changes after updating myAge value by *myAge

}
