package main

import "fmt"

func main() {
	car := new(Car) // pointer
	car2 := Car{}   // reference
	car3 := Car{
		Model: 2024,
		Color: "black",
		Engine: CarEngine{
			Version: 8,
		},
		Line : &Line{LineName: "trend line"},
	}
	//cars := make([]Car, 1) // array of cars

	fmt.Printf("%v", car)
	fmt.Println()
	fmt.Printf("%v", car2)
	fmt.Println()
	fmt.Printf("%v", car3)

	fmt.Println(r(car3)) // recibe una copia de car3 y solo la modifica en esa instancia

	fmt.Println(car3.Model) // comprobamos que no modifica el car3 original

	p(&car3) // recibe la direccion de memoria de car3, por lo tanto modifica el car3 original
	fmt.Println(car3.Model)
}

//Car is an exported structure
type Car struct {
	Model     int
	Color     string
	Engine    CarEngine
	Line      *Line // un puntero tiene una direccion de memoria como valor
	Insurance Insurance
}

type Line struct {
	LineName string
}

type CarEngine struct {
	Version int
}

type Insurance interface {
}

// recibe una copia, por lo tanto no modifica la original
func r(c Car) int {
	c.Model = 2001
	return c.Model
}

// recibe la direccion de memoria, por lo que modifica la original
func p(c *Car) {
	c.Model = 3001
}
