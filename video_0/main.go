package main

import "fmt"

// why go?
// https://yourbasic.org/golang/advantages-over-java-python/

func main(){
	dataTypesZeroValues()
}

func dataTypesZeroValues(){
	// data types
	var q int // zero value = 0
	var r float32 // zero value = 0 
	var s string // zero value = ""
	var t bool // zero value = false

	fmt.Printf("%v, %v, %v, %v", q, r, s, t)
}