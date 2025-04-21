package exported

import "fmt"

// funcion exportada (comienza con Mayuscula)
// con interface{} indicamos que se puede recibir valor de cualquier tipo de dato
func PrintConstant(c interface{}) {
	fmt.Printf("%v", c)
}

// funcion exportada que recibe varios argumentos (...) de cualquier tipo de dato, los agrupa en un slice de tipo interfaz vacia (interface{}) e imprime cada uno de ellos
func PrintConstants(constants... interface{}) {
	for _, c := range constants {
		fmt.Println(c)
	}
}