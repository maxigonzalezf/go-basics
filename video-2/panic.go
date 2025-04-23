package main

import (
	"fmt"
	"time"
)

func main() {
	//deferTest()
	fmt.Println("Hi from main")

	gWithRecover()
}

// defer es la ultima llamada en el cuerpo de la funcion
// sin importar donde este colocada la llamada
func deferTest() {
	defer timeTracker(time.Now(), "main")

	fmt.Println("first message")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("second message")
}

func timeTracker(t time.Time, name string) {
	elapsed := time.Since(t)

	fmt.Println(fmt.Sprintf("The function %s took %s", name, elapsed.String()))
}

// panic produce una interrupcion al ser llamado
// y hace el stackTrace de llamadas hacia arriba
// hasta encontrarse con quien genero esa goroutine (si es en main() se corta todo)
// nunca esta bueno usar panic por fuera del package main
// solo se debe usar para las funciones vitales para el software
func gWithPanic(i int) {
	// antes de ejecutar el panic, ejecuta el defer
	defer fmt.Println("not defer yet")
	fmt.Println(i)
	panic("going down dude")
}

// con recover evitamos que el programa se interrumpa con el panic
// (se puede usar con defer para asegurarse la ultima llamada)
func gWithRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r) // muestra el msg del panic (going down dude)
			fmt.Println("I'm still alive, recovered")
		}
	}()

	gWithPanic(99)
}
