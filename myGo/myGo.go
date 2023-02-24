package main

import (
	"fmt"
	"strings"
	"time"
)

var name, passkey, pasKey string

func main() {
	/*
			name
			passkey

		psuedo code
		if the password is pass, access granted
		else enter correct password



	*/
	fmt.Println("Welcome to registration page!")
	fmt.Println("Do provide your name ")
	fmt.Scanln(&name)
	fmt.Println("Enter your password")
	fmt.Scanln(&passkey)

	pasKey = strings.ToLower(passkey)
	if pasKey == "james@1234" {
		val := "Congratulations you made it!"
		vaL := strings.Repeat(val, 4)
		fmt.Printf("\n %s, %s", vaL, name)
	} else {

		fmt.Println("You have entered wrong password")
	}
	myTime()

}

func myTime() {
	t := time.Now()
	fmt.Printf("\n %.2d-%.2d-%.2d", t.Day(), t.Month(), t.Year())

}
