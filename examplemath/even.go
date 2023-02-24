package main

import "fmt"

var a  int

func main() {
	fmt.Println("Enter a number: ")
	fmt.Scanln(&a)
	if a%2 == 0 {
		fmt.Println(a, "is an even number ")
	} else if a%2 != 0 {

		fmt.Println(a, "Is an odd number")
	} else if a == 3 {
		fmt.Println(a, "are equivalent")
	}

}
