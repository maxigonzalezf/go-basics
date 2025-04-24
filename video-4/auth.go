package main

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JSON Web Token
// Estandar que define una forma compacta y autonoma de transmitir informacion de forma segura entre partes como un objeto JSON
// Son utiles en Autorizacion (por ej las act que realiza un usuario despues de iniciar sesion) e Intercambio de informacion
//  Estructura: consiste de tres partes separadas por puntos (header, payload, signature)
// xxxxx.yyyyy.zzzzz

// Header consiste en dos partes: el tipo del token, que es JWT y el algoritmo de firma utilizado, como RSA o HMAC SHA256
// ej: {
// "alg": "HS256"
// "typ": JWT
// }
// Luego, este JSON se codifica en Base64Url para formar la primera parte del JWT

// Payload (carga util)
// Contiene los claims (reclamos)
// Los claims son declaraciones sobre una entidad (gralmente el usuario) y datos adicionales
// Existen tres tipos de claims: registered, public y private
// Registered claims: conjunto de reclamos predefinidos que no son obligatorios, pero si recomendados. Son iss(emisor), exp(fecha de vencimiento), sub(asunto), aud(audiencia), etc
// Public claims: se pueden definir libremente, sin embargo, deben definirse en el Registro de JSON Web Tokens de la IANA
// Private claims: claims personalizadas creadas para compartir informacion entre partes que acuerdan usarlas y no son registered ni public
// ej: {
// "sub": "123456"
// "name" : "John Doe"
// "admin" : true
// }
// Luego, se codifica en Base64Url para formar la segunda parte del JWT

// Signature (firma)
// Para crear la parte de la firma, se debe tomar el Header codificado, el Payload codificado, un secret, el algoritmo especificado en el encabezado y firmarlo
// La firma se utiliza para verificar que el mensaje no se modifico durante el proceso
// En el caso de tokens firmados con una clave privada, tambien puede verificar que el remitente del JWT es quien dice ser
// ej:
// HMACSHA256(
//	 base64UrlEncode(header) + "." +
// 	 base64UrlEncode(payload),
//	 secret)

// TODO JUNTO
// La salida son tres Base64-URL strings separados por puntos que se pueden pasar facilmente en entornos HTML y HTTP
// Son mas compactas en comparacion con estandares basados en XML
// ej:
// eyJhbGci0iJIUzI1NiIsInR5cCI6IkpXVCJ9.
// eyJzdWIi0iIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4.
// 4pcPyMD09o1PSyXnrXCjTwXyr4BsezdI1AVTmud2fU4

const signKey = "changeMeInProduction"

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type JWTResponse struct {
	Token   string `json: "token"`
	Refresh string `json: "refresh_token"`
}

func (u *User) valid() bool {
	if u.Username == "maxi" && u.Password == "1234" {
		return true
	}
	return false
}

func getTokens() *JWTResponse {
	token := signToken()
	refresh := signRefreshToken()

	return &JWTResponse{
		Token:   token,
		Refresh: refresh,
	}
}

// firmar el token
func signToken() string {
	// utilizamos una libreria que eligio tomas para el ejemplo para generar el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"exp": time.Date(2025, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	t, err := token.SignedString(signKey) // tal vez es []byte(signKey)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return t
}

// hacer refresh del token
func signRefreshToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "baz",
		// un mes despues por ejemplo
		"exp":     time.Date(2025, 11, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"refresh": true,
	})
	t, err := token.SignedString(signKey) // tal vez es []byte(signKey)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return t
}
