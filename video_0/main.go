package main

import "fmt"

// why go?
// https://yourbasic.org/golang/advantages-over-java-python/

// Constants
const myConst int = 7
const myOtherConst = "String constante"
const (
	agrupatedConst  = 3
	agrupatedConst2 = 2
)

func main() {
	// this is a one line comment
	/*
		This is
		a multiline comment
	*/
	dataTypesZeroValues()
	dataTypesWithValues()
	dataTypesGroup()
	printConstants()
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

func printConstants() {
	fmt.Printf("%v, %v, %v, %v", myConst, myOtherConst,
				agrupatedConst, agrupatedConst2)
}
