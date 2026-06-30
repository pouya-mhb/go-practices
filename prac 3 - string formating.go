package main

import "fmt"

func main () {
	var name string
	name = "pouya"

	lastName := "MHB"

	fmt.Println (name + lastName)

	fullName := name + lastName
	fmt.Println (fullName)

	fmt.Printf("Hi my name is %v",fullName)

	fullName2:= fmt.Sprintf("%v %v", name, lastName)
	fmt.Println (fullName2)
}