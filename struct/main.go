package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	firstName string
	lastName  string
	birthDay  string
	created   time.Time
}

func newUser(fname string, lname string, bDay string) User {

	created := time.Now()

	var userData User
	userData = User{
		firstName: fname,
		lastName:  lname,
		birthDay:  bDay,
		created:   created,
	}

	return userData

}

func main() {

	firstname := getUserData("enter name : ")
	lastname := getUserData("enter last name : ")
	birthday := getUserData("enter birthday : ")

	userData := newUser(firstname, lastname, birthday)
	fmt.Println(userData)
}

func getUserData(text string) string {
	fmt.Print(text)
	userInput, _ := reader.ReadString('\n')
	cleanedInput := strings.TrimSpace(userInput)
	return cleanedInput
}
