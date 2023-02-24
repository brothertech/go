package main
import "fmt"
//var i, a int

func main() {
//myTriangle()
//aphaShape()
diamond()
}
/*
func myTriangle(){

	
	//this code will print pyramind
	for i=3; i<=10; i++ {
		for c := 0; c < i; c++ {

			fmt.Print("##")

			for d:=5; d<c; d++{
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
		
	}

	
}

//this function will print A-D in triangle instead of using #
func aphaShape(){
for i:='A'; i<='Z'; i++{
	for a:='A'; a<=i; a++{
		fmt.Printf("%c ", a)
	}
	fmt.Println()
}
		
}
*/

//function from here will print a diamond shape with numbers

func diamond(){
num := 1
rows := 45
space := rows/2

for i:=1; i<=rows; i++{
	for j:=1; j<=space; j++{
		fmt.Printf(" ")
	}
	count := num/2 +1

	for k := 1; k<=num; k++ {

		fmt.Printf("%d", count)
		if k <=num/2{
			count--
		} else{
			count++
		}
	}
	fmt.Println()

	if i<=rows/2{
		space =space-1
		num = num+2
	} else{
		space =space+1
		num = num-2
	}

}



}

