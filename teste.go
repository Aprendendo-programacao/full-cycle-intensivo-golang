package main

import (
	"errors"
	"fmt"
)

func main() {
	a, b, err := nome("wesley", 10)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(a, b)

	anonymousFunction := func() int {
		return 1
	}

	fmt.Println(anonymousFunction())
}

func nome(a string, b int) (string, int, error) {

	if b > 10 {
		return "", 0, errors.New("b precisa ser menor ou igual a 10")
	}

	return a, b, nil
}