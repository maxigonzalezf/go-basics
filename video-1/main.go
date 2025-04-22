package main

import "fmt"

// Interfaces
// Son implicitas
// struct --> interface (las estructuras implementan las interfaces)

func main () {
	fmt.Println("I want a coffee...")

	var (
		i ItalianCoffeeMachine
		c ColombianCoffeeMachine
	)

	// se debe pasar la direccion de memoria de la estructura que implementa la interfaz (las interfaces son punteros)
	italianCoffee := GetCoffee(&i, 10)

	italianCoffee.PrintCoffee()

	colombianCoffee := GetCoffee(&c, 25)

	colombianCoffee.PrintCoffee()

	//...

	machine := CoffeeMachine{&c}

	machineCCoffee := machine.MakeCoffee(35)

	machineCCoffee.PrintCoffee()

	//...

	supreme := Supreme{}
	supremeCoffee, status := SupremeCaller(&supreme)

	supremeCoffee.PrintCoffee()
	fmt.Println("status: " + status)
}

type Coffee struct {
	Intensity int
	Region string
}

func (c *Coffee) PrintCoffee() {
	fmt.Println(fmt.Sprintf("This coffee is from %s and intensity is %d", c.Region, c.Intensity))
}

// CoffeeMaker
type CoffeeMaker interface {
	MakeCoffee(intensity int) Coffee
}

type ItalianCoffeeMachine struct {
}

type ColombianCoffeeMachine struct {
}

// Metodo implementado por la estructura ItalianCoffeeMachine
func (i *ItalianCoffeeMachine) MakeCoffee(intensity int) Coffee {
	return Coffee{Intensity: intensity, Region: "Italy"}
}

func (c *ColombianCoffeeMachine) MakeCoffee(intensity int) Coffee {
	return Coffee{Intensity: intensity, Region: "Colombia"}
}

// .....
// Metodo polimorfico (acepta cualquier estructura que implementa la interfaz CoffeeMaker)
func GetCoffee(coffeeMaker CoffeeMaker, i int) Coffee {
	return coffeeMaker.MakeCoffee(i)
}

type CoffeeMachine struct {
	CoffeeMaker
}

type Supreme struct {}

type SupremeMachine interface {
	CoffeeMaker
	CheckMachine() string
}

func (s *Supreme) MakeCoffee(i int) Coffee {
	return Coffee{Intensity: i, Region: "Unknown"}
}

func (s *Supreme) CheckMachine() string {
	return "looks good!"
}

func SupremeCaller(sm SupremeMachine) (coffee Coffee, status string) {
	coffee = sm.MakeCoffee(99)
	status = sm.CheckMachine()

	return
}