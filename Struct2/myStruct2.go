package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)

}

func main() {
	//1-struct as a value type:
	var pers1 Person
	pers1.firstName = "Abiodun"
	pers1.lastName = "Oguntuyi"
	upPerson(&pers1)
	fmt.Println("This is value type of struct")
	fmt.Printf("Hi .... %s %s, how are you doing today? \n", pers1.firstName, pers1.lastName)
//2-struct as a pointer type

var pers2 =new(Person)
pers2.firstName = "Abiodun"
pers2.lastName = "Oguntuyi"
upPerson(pers2)
fmt.Println("This is struct as pointer type")
fmt.Printf("Hi .... %s %s, how are you doing today? \n", pers2.firstName, pers2.lastName)

//3-struct as a literal:
pers3 := &Person{"Abiodun", "Oguntuyi"}
upPerson(pers3)
fmt.Printf("Hi... %s %s", pers3.firstName, pers3.lastName)


}
