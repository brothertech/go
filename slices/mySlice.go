package main

import "fmt"

func main() {
	mySlice1 := []int{2, 3, 4, 5, 6}

	fmt.Printf("The capacity of the slice is %d\n", cap(mySlice1))
	fmt.Printf("The lenght of the array is %d", len(mySlice1))
	fmt.Println("The arrray is ", mySlice1)
	mySlice2 := append(mySlice1, 3, 9, 5, 6, 7)
	fmt.Printf("The length of the new slice is %d\n", len(mySlice2))
	fmt.Printf("The capacity of the slice is %d\n", cap(mySlice2))
	fmt.Println("The new slice is ", mySlice2)
	fmt.Println("The combination of the two slices gives ", append(mySlice1, mySlice2...))
	fmt.Printf("The capacityo of the two combined is %d\n", cap(append(mySlice1, mySlice2...)))

	fmt.Printf(" And the lenght is %d\n", len(append(mySlice1, mySlice2...)))
}
