package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	err := terror()

	if err != nil {
		panic("time is not even")
	}

	// inline if con handling de error
	if _, err := (division(3, 0)); err != nil {
		fmt.Println(err.Error())
		return
	}

}

// las interfaces en golang son punteros

func terror() error { // error es una interfaz
	if time.Now().Unix()%2 == 0 {
		return nil
	}
	return &MyError{Msg: "This is an error"}
}

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func division(a, b int) (float32, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide by zero")
	}

	return float32(a / b), nil
}
