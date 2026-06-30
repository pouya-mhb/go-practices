package main

import "fmt"
import "time"

func main () {

	var name string
	name = "pouya"

	lastName := "MHB"

	fmt.Println (name + lastName)

	today := time.Now().Year()
	birthday := 1997
	age :=  today - birthday

	fmt.Println ("age : ", age)

	today2 := time.Now()
	nextYear := today2 + today2
	fmt.Println (nextYear)
}