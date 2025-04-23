package main

import "fmt"

type fn func(i int) int // ver func fnTypes()

func main() {
	strValue := fnArgs(10, func(i int) bool {
		return i % 2 == 0
	})

	fmt.Println(strValue)
}

// palabra reservada 'func'
// name -> f1
// args -> () -- opcional
// tipo de dato de return -- opcional (void return)
func f1() {
	fmt.Println("No args and no return")
}

func f2(i int) int {
	return i + 10
}

// return multiple (int, string, error)
func f3(i, j int) (int, int) {
	return i + j, i * j
}

// return nombrado (utilizado con naked return)
func f4(i, j int) (sum int, multip int) {
	sum = i + j
	multip = i * j
	return // naked return
	//return sum, multip // multi named return (valid)
}

// pueden recibir de 0 a n argumentos
func invariantsArgs(i ...int) (sum int) {
	for _, value := range i {
		sum += value
	}

	return
}

func anonFunctions() {
	fn := func(i int) int {
		return i + 10
	}

	fmt.Printf("%T", fn) // de tipo func(int) int

	fg := func(s string) int {
		return 0
	}

	fmt.Printf("%T", fg) // de tipo func(string) int

	fn(8)
	fg("100")
}

// Se define un tipo fn (arriba antes de la func main()) para asignarsela a una variable y la defino en la llamada
func fnTypes() {
	var f fn

	f = func(i int) int {
		return 0
	}

	f(0)
}

// fnArgs recibe por parametro una funcion anonima que definiremos al momento de la llamada
// (en este caso retorna un booleano que dice si el numero recibido por parametro es par o impar)
func fnArgs(n int, evenFn func(i int) bool) string {

	if evenFn(n) {
		return "yes"
	} else {
		return "no"
	}
}