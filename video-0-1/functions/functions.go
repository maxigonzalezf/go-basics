package main

import "fmt"

func main() {
	unexported(2, "hola", func(u int){
		fmt.Println(u)
	}, "hello", "world")

	fmt.Println(withReturnOneTypeValue())
	fmt.Println(withReturnManyTypesValues())
	a, b := withReturnManyTypesValuesAndNames()
	fmt.Printf("%d - %s", a, b)

}

// unexported cualquier tipo puede funcionar como argumento, incluso funciones
// invariant args son 0...n, la sintaxis es [name ...(type)]
func unexported(i int, s string, fn func(u int), t ...string) {
	fmt.Println(i)
	fmt.Println(s)

	fn(8)

	for a, b := range t {
		fmt.Printf("%d - %s", a, b)
		fmt.Println()
	}
}

func withReturnOneTypeValue() int {
	return 1
}

func withReturnManyTypesValues() (int, string) {
	return 3, "golang"
}

func withReturnManyTypesValuesAndNames() (i int, s string) {
	i = 10
	s = "naked return"

	return // no hace falta especificar ya que entiende que tiene que devolver las variables i y s (si no tienen valor asignado, les asigna el zero value)
}