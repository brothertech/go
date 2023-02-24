package main
import "fmt"

var first, second string
func main(){
		fmt.Println("Enter your first name: ")
		fmt.Scanln(&first)
		fmt.Println("Enter your second name: ")
		fmt.Scanln(&second)
		if ((first == "Oguntuyi") && (second== "Abiodun")){
			fmt.Println("You are welcome!")

		}else{

			fmt.Println("Access not granted!")
		}
		fmt.Println("Your name is ", first, second)





}