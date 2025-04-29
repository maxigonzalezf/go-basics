package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// common errors in golang

// 1. Declarar constantes con iota y sin valor por defecto
type status int

const (
	// la idea es no dejarlo sin tipo
	statusUnknown  status = iota // 0 la idea es mapearlo al zero value para poder trabajar los demas datos con valores conocidos
	statusActive                 // 1
	_                            // skip number 2
	statusInactive               // 3
)

// 2. No usar timeouts en llamadas http
func httpCall() {
	// Siempre rapeada en un contexto que se pueda cancelar (en este caso despues de 5 segundos)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// error nro 3 (ignorar errores)
	req, _ := http.NewRequestWithContext(ctx, "POST", "google.com", nil)

	client := http.Client{}
	client.Do(req)

	// como servidor, se configura el readtimeout y writetimeout
	serv := http.Server{
		Addr:              "",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       time.Second * 10, // timeout
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    1024,
	}
	serv.ListenAndServe()
}

// 3. Ignorar errores, o manejarlos desde una misma funcion
func ignoreErr() {
	f, err := os.Create("this.txt")

	if err != nil {
		// log Fatal no, es igual que un panic
		// siempre se debe loguear el error
		log.Print("the file cannot be created " + err.Error()) // en caso de devolver el error, se deja de loguearlo, se envia a la capa de presentacion y se transforma en un error que le sirva al cliente
		return                                                 // cortar la funcion en caso de haber una funcion
	}

	// este tipo de errores se puede "perdonar" ignorarlos (no cerrar el archivo)
	defer f.Close() // uso el defer para manejar el error en caso que haya e igualmente cierre el archivo
}

// 4. Muchas interfaces
// La idea de las interfaces es que donde las definamos sepamos que las vamos a usar
// Usarlas internamente, y las que vamos a externalizar, darle la posiblidad al cliente que sea él el que pueda implementarla y que defina el modo de usarla

// 5. Cuando ocurre un error, recordar cerrar, cancelar, etc los componentes usados
// Es muy importante cerrar las conexiones entre nuestro runtime y el sistema operativo (utilizar el defer (como en los ejemplos 2 y 3) para asegurar que sucedan)
// Golang tiene el garbage collector que ayuda a liberar memoria (eliminar las variables, etc no utilizadas que quedan por ahi)

// 6. Si usamos read/write lock, separar bien lectura y escritura, ademas de liberar los locks
func usingLocks() {
	l := sync.Mutex{}

	l.Lock()
	fmt.Print("Locked")
	defer l.Unlock() // nunca olvidarse de liberar los locks
}

func usingRWLocks() {
	l := sync.RWMutex{} // mutual exclusion

	// writing
	l.Lock()
	fmt.Print("writing")
	defer l.Unlock() // defer por si hay otros flujos de salida en la funcion y olvidarme de hacer el unlock despues(asegurar que se libere el lock)

	// reading
	l.RLock()
	fmt.Print("reading")
	l.RUnlock()
}

// 7. No test unitarios
type PatientStorage interface {
	Save(name string)
}

type PatientStg struct{}

func (p *PatientStg) Save(name string) {
	// SQL Calls
	// logs...
}

// Se puede mockear lo que esta detras de una interfaz
type MockPatientStg struct {
	// mock
}

func (p *MockPatientStg) Save(name string) {
	// SQL Calls
	// logs...

	if name == "fail" {
		panic("")
	}
}

func main() {

}
