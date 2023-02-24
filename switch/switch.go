package main

import "fmt"

var grade float32
var result bool
var name string

//var result float64
func main() {
	fmt.Println("Enter your grade!")
	fmt.Scanln(&grade)

	fmt.Println("Welcome enter your name")
	fmt.Scanln(&name)

	switch name {
	case "Abiodun":
		result = grade > 70 || grade <= 100
		fmt.Println("Wow! you got an A", grade)

	//result=	(grade > 70)
	//fmt.Println("Wow! you got %f and that has earned you an A", result)

	case "Sham":
		result = grade <= 60.0 || grade >= 69.0
		fmt.Println("You performed well and that has earned you a B", grade)

	case "Joe":
		result = grade<=50 || grade>=59
		fmt.Printf("That's not bad %s, you got C", name)
	default:
		fmt.Printf("Oops! seems you didn't take the examination")
	}


}
