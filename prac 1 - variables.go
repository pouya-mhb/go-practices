package main

import "fmt"

func main(){

	// var txt string
	// txt = "This is a text"

	// var txt string = "This is a text"

	txt := "this is a text"


	fmt.Println(txt)

	var newRune rune
	newRune = '$'
	fmt.Println (newRune)
	fmt.Println (string (newRune))

	var newByte byte
	newByte = 'a'
	fmt.Println (newByte)
	fmt.Println (string (newByte))
}