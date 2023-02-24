package main

import (
	"fmt"
)

func main() {
	//this example is looking at map functions
	var value int
	var isPresent bool
	map1 := make(map[string]int)
	//map2 :=make(map[string]int)

	map1["Nigeria"] = 50
	map1["Algeria"] = 40
	map1["Washinton"] = 65
	value, isPresent = map1["Washinton"]
	if isPresent {
		fmt.Printf("Is Washinton present in the map 1? %t, is there present and its value is %d \n", isPresent, value)

	} else {
		fmt.Printf("The input seems to be incorrect, please enter correct inout and try again \n")

	}

	map1["Paris"] = 64

	delete(map1, "Paris")
	value, isPresent = map1["Paris"]
	if isPresent {

		fmt.Printf("\nParis is my dream country!")

	} else {
		fmt.Printf("The input is incorrect! because the response is %t \n", isPresent)

	}

	fmt.Printf("the value of Washintoin in the map is %d ", value)
	fmt.Printf(" ")
	fmt.Printf(" ")
	fmt.Printf(" ")
	StateCapital()
	fmt.Println(" ")
	dayInWeek()
	fmt.Println(" ")
	definedMap()
}

func definedMap() {

	map1 := make(map[int]float32)

	map1[1] = 10
	map1[2] = 20
	map1[3] = 30
	map1[4] = 40

	for key, value := range map1 {

		for i := 0; i <= len(map1); i++ {

			fmt.Print(i)
			fmt.Printf("The values and keys are %d %f \n", key, value)

		}

	}
//fft
}

// a function containing states in Nigeria and their capitals
func StateCapital() {
	map1 := map[string]string{"Ondo": "Akure", "Kaduna": "Kaduna", "Benue": "Markurdi", "Ekiti": "Ado Ekiti"}
	for key, value := range map1 {

		fmt.Println(key, " is a state in Nigeria and its metropolis is ", value)
	}

}

func dayInWeek() {
	map1 := map[string]int{"Monday": 2, "Tuesday": 3, "Wednesday": 4, "Thursday": 5, "Friday": 6, "Saturday": 7, "Sunday": 8}

	for key, value := range map1 {
		fmt.Println(key, " belong to number ", value, "of the week or how do you see it? ")

	}

}
