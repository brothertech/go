package main

import (
	//"rand"
	"fmt"
	//"rand"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to new world of programming guys!")
	/*
			for i := 0; i <= 10; i++ {
				for c := 0; c < i; c++ {
					fmt.Print("**")

				}
				could not import rand (cannot find package "rand" in any of
		C:\Program Files\Go\src\rand (from $GOROOT)
		C:\Users\ENGR ABIODUN\go\src\rand (from $GOPATH))compilerBrokenImport
				for k := 5; k < c; k++ {

					fmt.Print("*")
				}

			}


	*/
	String()
	myTriangle()
	mySlice1()
	diamond()
	aphaShape()
	stringCase()
}

/*
func myTriangle() {

	//this code will print pyramind
	for i = 3; i <= 10; i++ {
		for c := 0; c < i; c++ {

			fmt.Print("##")

			for d := 5; d < c; d++ {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")

	}

}
*/
//Another function begins here
func diamond() {
	num := 1
	rows := 10
	space := rows / 2

	for i := 1; i <= rows; i++ {
		for j := 1; j <= space; j++ {
			fmt.Printf(" ")
		}
		count := num/2 + 1

		for k := 1; k <= num; k++ {

			fmt.Printf("%d", count)
			if k <= num/2 {
				count--
			} else {
				count++
			}
		}
		fmt.Println()

		if i <= rows/2 {
			space = space - 1
			num = num + 2
		} else {
			space = space + 1
			num = num - 2
		}

	}

}
func aphaShape() {
	for i := 'A'; i <= 'Z'; i++ {
		for k := 'A'; k < i; k++ {
			fmt.Printf("%c", k)

		}
		fmt.Println()

	}

}

// this function is centered on slice and array

func mySlice1() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d /\n", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / \n", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / \n ", 100*rand.Float32())
	}
}

//Using Uppercase and Lowercase

func stringCase() {
	var lower string
	var upper string
	orig := "How are you doing today!"
	lower = strings.ToLower(orig)
	upper = strings.ToUpper(orig)

	//I want to apply trim here
	space := "    Hello     "
	space1 := "    Hello world    "
	// Using Trim() function
	res1 := strings.Trim(space, "!")
	res2 := strings.Trim(space1, "@$")

	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1, len(space), "lenght after trimming is ", len(res1))
	fmt.Println("Result 2:", res2, len(space1), "lenght after trimming is ", len(res2))
	fmt.Printf("This is showing lower -> %s \n while this is showing upper-> %s ", lower, upper)

}

func myTriangle() {
	for a := 0; a < 10; a++ {
		for b := 2; b < a; b++ {
			fmt.Print("G")
		}
		fmt.Println()
	}

}
func String() {
	var str = "Hello Mr Gideon"
	stri1 := strings.Repeat(str, 3)
	stril2 := strings.Replace(str, "Hello Mr Gideon", "I am Engr Abiodun", 3) //this is for replace
	stril3 := strings.Repeat(stril2, 2)                                       // repeat string
	fmt.Println("This string is printing 3 times", stri1, " after replace we get ->", stril3)

}
