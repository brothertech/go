package main

import (
	"fmt"
	"strings"
)

type Address struct {
	Address    string
	name       string
	birth_date string
	photo      string
}

func upAddress(a *Address) {
	a.Address = strings.ToLower(a.Address)
	a.name = strings.ToUpper(a.name)
	a.birth_date = strings.ToLower(a.birth_date)
	a.photo = strings.ToUpper(a.photo)

}
func main() {
	//1- using value type

	var my Address
	my.Address = "Iselu Owo, Ondo State, Nigeria"
	my.name = "Abiodun Oguntuyi"
	my.birth_date = "2023"
	my.photo = "loading"
	upAddress(&my)
	fmt.Printf("%s is living at %s and was born in %s while his photo is still %s", my.name, my.Address, my.birth_date, my.photo)
	fmt.Println("==================")
	//using as pointer type
	var my2 = new(Address)
	my.Address = "Iselu Owo, Ondo State, Nigeria"
	my2.name = "Abiodun Oguntuyi"
	my2.birth_date = "2023"
	my2.photo = "loading"
	upAddress(my2)
	fmt.Printf("%s is living at %s and was born in %s while his photo is still %s", my2.name, my2.Address, my2.birth_date, my2.photo)

	fmt.Println("==================")
	//3-as literal
	my3 := &Address{"Abiodun Oguntuyi", "Iselu Owo, Ondo State, Nigeria", "2023", "loading"}
	upAddress(my3)
	fmt.Printf("%s is living at %s and was born in %s while his photo is still %s", my3.name, my3.Address, my3.birth_date, my3.photo)

}
