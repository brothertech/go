package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&b)
	if a > b {
		fmt.Print(a, "is greater than ", b)
		//fmt.Printf("value of a: %d\n", a)
	}

}
