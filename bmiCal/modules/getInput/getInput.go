package getinput

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func GetInput() (string, string) {

	fmt.Println("Enter height : ")
	height, _ := reader.ReadString('\n')

	fmt.Println("Enter weight : ")
	weight, _ := reader.ReadString('\n')

	return height, weight

}
