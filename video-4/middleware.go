package main

// El cliente realiza una request http a nuestro servidor Web
// El servidor realiza una response http al cliente
// El middleware se encarga de pre-procesar esa request antes de llegar al handler que genera la response
// La idea es que el middleware sea una funcion que se pueda compartir entre varios handler
// por eso deben ser lo mas genericos posibles (no tanta logica de negocio)
// generalmente no tiene tanto sentido hacer un middleware para un solo endpoint

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	server := newWebServer("9999")
	routes()
	server.start()

}

type webServer struct {
	s *http.Server
}

// funcion para crear puntero del webServer
// recibe el port
func newWebServer(port string) *webServer {
	s := &http.Server{
		Addr: ":" + port,
	}

	return &webServer{
		s: s,
	}
}

// funcion para arrancar el webServer
func (w *webServer) start() {
	log.Println("server start at port" + w.s.Addr)
	log.Fatal(w.s.ListenAndServe())
}

//  rutas de http requests/responses
func routes() {
	// localhost.com/9999/test nos devuelve la hora actual
	// ejemplo sin middleware
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().String()
		_, _ = w.Write([]byte(fmt.Sprintf("the time is %s", &t)))
	})

	// ejemplo con middleware
	http.HandleFunc("/hello", middleware(hello()))
}

func hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello world"))
	}
}

// una funcion middleware recibe una interfaz Handler y devuelve una HandleFunc
func middleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// logica del middleware
		m := r.Method
		log.Println(m)
		log.Println(r.RemoteAddr)
		// se llama al Handler con ServeHTTP
		next.ServeHTTP(w, r)
	}
}