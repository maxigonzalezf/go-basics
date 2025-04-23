package main

import "fmt"

func main() {
	// asigno la funcion a una variable 'fn'
	fn := rFn("hello")
	
	//fmt.Printf("%T", fn) // func()

	// llamo a la funcion a traves de la variable 'fn'
	fn()

	// asigno la funcion test a la variable fnOperation
	fnOperation := test(1, 2, operation) // operation es de tipo funcion con la misma firma (recibe dos int y retorna un int)
	// fnOperation2 := test(1, 2, operationExample2()) // se llama a la funcion operationExample2() que es del tipo funcion con la misma firma

	// asigno a la variable result
	// el int retornado por la funcion test asignada a la variable fnOperation
	result := fnOperation()
	// result2 := fnOperation2()

	fmt.Printf("The result of the operation is: %d", result)
}

// funcion que retorna otra funcion
func rFn(s string) func() {
	fmt.Println(s)

	return func() {
		fmt.Println("Inside the return " + s)
	}
}
// Es util para hacer un post procesamiento de alguna data

// funcion que recibe una funcion y retorna otra funcion
func test(i, j int, fn1 func(i, j int) int) func() int {
	res := fn1(i,j)
	return func() int {
		return res * 100
	}
}

func operation (i, j int) int {
	return (i + 150) * (j + 200)
}

func operationExample2() func(int, int) int {
	return func(i, j int) int {
		return (i + 150) * (j + 200)
	}
}