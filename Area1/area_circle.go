package main

import "fmt"

const PI = 3.142

func main() {
	var area, p, cir, sec float64
	var theta, radius = 40.0, 7.0
	p = theta / 360.0
	area = PI * radius * radius
	cir = 2 * PI * radius
	sec = PI * p * radius * radius

	fmt.Println("The area of the circle is \n", area, "cm^2")
	fmt.Println("The circumference of the circle is:\n", cir, "cm")
	fmt.Println("The sector of the circle is:%f ", sec)

}
