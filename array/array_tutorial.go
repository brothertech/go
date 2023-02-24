package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [4]string{"Joy", "Abiodun", "Michele", "Kenny"}

	//arr3 := [4]int{}
	/*
		for i := 0; i <= 4; i++ {
			fmt.Print("Enter the value one by one: ")
			fmt.Scanln(&arr3)
		}
		fmt.Println(arr3)
	*/
	fmt.Println(arr1, " ", arr2)
	var a, b = arr1[1], arr1[0]
	v := a + b
	fmt.Printf("Adding the array index 1 and 0 gives %d\n", v)
	fmt.Println("Lenght of the first array (arr1) is ", len(arr1), " and arr2 is", len(arr2))
	fmt.Printf("The capacity of the array is %d", cap(arr1)) // cap will show the capacity of an array
}
