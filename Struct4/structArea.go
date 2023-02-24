package main

type Area struct {
	area int
	l    int
	br   int
}

func myArea(a *Area) {
	a.br = 5
	a.l = 2
	a.area = a.br * a.l

}

func main() {
	var my = Area
	ar := new(Area)
	//ar.l = 5
	//ar.br = 7
	//ar.area = ar.l * ar.br

}
