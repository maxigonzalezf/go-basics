package main

// El cliente realiza una request http a nuestro servidor Web
// El servidor realiza una response http al cliente
// El middleware se encarga de pre-procesar esa request antes de llegar al handler que genera la response
// La idea es que el middleware sea una funcion que se pueda compartir entre varios handler
// por eso deben ser lo mas genericos posibles (no tanta logica de negocio)
// generalmente no tiene tanto sentido hacer un middleware para un solo endpoint

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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

// rutas de http requests/responses
func routes() {
	// localhost.com/9999/test nos devuelve la hora actual
	// ejemplo sin middleware
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().String()
		_, _ = w.Write([]byte(fmt.Sprintf("the time is %s", &t)))
	})

	// ejemplo con middleware
	http.HandleFunc("/hello", middleware(hello()))

	http.HandleFunc("/login", userHandler())

	http.HandleFunc("/greeting", validateTokenMiddleware(secureGreeting()))
}

// endpoint seguro, validado por el middleware
func secureGreeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from secure endpoint"))
	}
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

func userHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// formato de respuesta json ({token: ..., refresh_token:...})
		writer.Header().Set("Content-Type", "Application/json")
		if request.Method == "POST" {
			body := request.Body // array de bytes para mantener strings en el flujo de trabajo
			defer body.Close()   // cerrar el body para ahorrarle tiempo al runtime de Go
			var user User
			// mapear el JSON a la estructura User
			err := json.NewDecoder(body).Decode(&user)

			// manejamos un posible error
			if err != nil {
				http.Error(writer, "cannot decode json", http.StatusBadRequest)
				return
			}

			if user.valid() {
				res := getTokens()
				// en este ejemplo ignoramos el posible error
				_ = json.NewEncoder(writer).Encode(&res)
			} else {
				http.Error(writer, "bad credentials", http.StatusUnauthorized)
				return
			}
		} else {
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// middleware para validar los token
// la idea es que si falla la autorizacion por estar el token vencido, el cliente pueda utilizar el refresh_token
// ya sea enviando ambos token el cliente, o negociar otro par de tokens nuevos con el refresh_token
func validateTokenMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// obtiene el token del header de Authorization
		bearerToken := r.Header.Get("Authorization")
		// chequea que sea valido
		if validateToken(bearerToken) {
			next.ServeHTTP(w, r)
			return
		}

		http.Error(w, "invalid token", http.StatusUnauthorized)
	}
}

// funcion para validar el token obtenido del Header de Authorization
func validateToken(s string) bool {
	// el value de Authorization es bearer (...) siendo (...) el token
	// por lo tanto le saca el 'bearer ' al string recibido y solo deja el token
	token := strings.Replace(s, "bearer ", "", 1) // bearer ignore case seria lo correcto
	mySigningKey := []byte("AllYourBase")
	// validar el algoritmo
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		//mySigningKey es un []byte que contiene el secret, ej []byte("my_secret_key")
		return mySigningKey, nil
	})

	if err != nil {
		log.Println(err.Error())
		return false
	}

	// t se vuelve a validar con su propia validacion
	return t.Valid
}
