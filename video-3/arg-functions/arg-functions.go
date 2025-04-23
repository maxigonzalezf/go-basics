package main

import (
	"errors"
	"fmt"
)

type argFNInt func(i, j int) (string, error)

func main() {
	s, err := fh(func(i, j int) (string, error) {
		if i > j {
			return "i es mayor", nil
		}

		return "", errors.New("j no puede ser mayor")
	})

	fmt.Printf("%v, %v", s, err)
}

func fh(a argFNInt) (string, error) {
	s, err := a(1, 2)

	if err != nil {
		return "", err
	}

	return s, nil
}
