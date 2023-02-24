package main

import "fmt"

func math(){
	var num1 int = 1
	var num2 int = 2

	sum := num1 + num2
	product := num1*num2

	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)
	fmt.Printf("The multiplication of %d and %d is %d\n", num1, num2, product)
}
