package main

import (
	"fmt"

	"github.com/maxigonzalezf/go-tomaslingotti/video_0/exported"
)

// why go?
// https://yourbasic.org/golang/advantages-over-java-python/

// unexported & exported
// unexported se pueden llamar desde distintos archivos dentro de un mismo paquete
// exported se pueden llamar desde archivos de otros paquetes

// Constants
const myConst int = 7
const myOtherConst = 8
const (
	agrupatedConst  = "Holis"
	agrupatedConst2 = 2
)

// entry point
func main() {
	// this is a one line comment
	/*
		This is
		a multiline comment
	*/
	dataTypesZeroValues()
	println("")
	dataTypesWithValues()
	println("")
	dataTypesGroup()
	println("")
	printInt(myConst)                      // utilizo funcion de otro archivo dentro del mismo package
	exported.PrintConstant(agrupatedConst) // funcion importada del package exported
	println("")
	exported.PrintConstants(myConst, myOtherConst, agrupatedConst, agrupatedConst2)
	otherTypeOfAssignation()
	myOtherTypes()
	otherArrayTypes()
}

func dataTypesZeroValues() {
	// data types
	var q int     // zero value = 0
	var r float32 // zero value = 0
	var s string  // zero value = ""
	var t bool    // zero value = false

	fmt.Printf("%v, %v, %v, %v", q, r, s, t)
}

func dataTypesWithValues() {
	// data types
	var q int = 0
	var r float32 = 0.32
	var f = 0.64           // float64 por defecto
	var s string = "hello" // same as var 's = "hello"'
	var t bool = true

	fmt.Printf("%v, %v, %v, %v, %v", q, r, f, s, t)
}

func dataTypesGroup() {
	var (
		q    = 1
		name = "maxi"
		b    = true
	)
	fmt.Printf("%v, %v, %v", q, name, b)
}

func otherTypeOfAssignation() {
	myNumber := 8
	fmt.Println(myNumber)
	myString := "my name is Maxi"
	fmt.Println(myString)
}

func myOtherTypes() {
	// array []int = {1, 2, 3, 4} // INVALID OPERATION
	arr := []int{1, 2, 3, 4, 5}

	// for loop convencional
	for i := 0; i < 10; i++ {
		// ...
		break
	}

	// for mas indicado para recorrer arrays
	for index, value := range arr {
		fmt.Printf("%v) %v", index, value)
		fmt.Println()
	}

	// len() devuelve el largo del arreglo
	arrLen := len(arr)
	fmt.Println(arrLen)
}

func otherArrayTypes() {
	var sl []int // nil
	sl = []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v", sl)
	// la funcion append agrega el 6 y devuelve el nuevo slice con el valor agregado
	sl = append(sl, 6)
	fmt.Printf("Slice: %v", sl)
}
