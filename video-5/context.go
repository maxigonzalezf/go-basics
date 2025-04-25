package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// testGetContest()
	fmt.Println("server started!")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/greeting", middleware(greetingHandler))
	http.ListenAndServe(":9191", nil)
}

// Context es una interfaz
// Enviar y recibir valores es una de sus principales funciones
// Enviar señales es otra funcion

func getContext() context.Context {
	ctx := context.Background() // el limite es el main() que en este caso es el holder
	// ctxDone := context.TODO() // generalmente se usa para dar contexto en tests. Siempre esta vacio (no nil)
	ctxWithValue := context.WithValue(ctx, 1, 999)
	return ctxWithValue
}

func testGetContest() {
	ctx := getContext()
	value := ctx.Value(1) // con la key se extrae el valor
	fmt.Println(value)
}

// la Request tiene un contexto muy importante que es el ciclo de vida de la llamada a nuestra web server
// comienza desde que llega a nuestra web server, pasando por toda la logica de negocios hasta que termina con un status http
// se devuelve al cliente y termina el ciclo de vida de la request
// en ese ciclo de vida se pueden establecer contextos, para enviar señales y enviar y recibir valores
// Si usamos un middleware, dentro del cual se puede ir propagando ese contexto por otro middleware o a la funcion Handler y transportando data de ser necesario
// (ejemplo: TraceID) se genera un string aleatorio

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // contexto de la request
	var err error
	defer fmt.Println("after signal interrupt")

	select {
	case <-time.After(10 * time.Second):
	case <-ctx.Done():
		err = ctx.Err()
	}

	if err != nil {
		fmt.Printf("error is %s", err.Error())
	}
}

// ejemplo de contexto con middleware
func middleware(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "traceID", "als2d3f4g5")
		next.ServeHTTP(w, r.WithContext(ctx)) // debemos pasar la request con el contexto
	}
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	traceID := r.Context().Value("traceID")
	w.Write([]byte(traceID.(string)))
}
