package main

import (
	"fmt"
	"net/http"
)

const xpto int = 1
const nome string = "wes"

const (
	a string = "aaa"
	b string = "bbb"
	c string = "ccc"
)

func main() {

	response, _ := http.Get("https://www.google.com")

	fmt.Println(response.Body)

}
