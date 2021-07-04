package main

import "fmt"

func main() {
	x := 10
	result := getPointerValue(&x)
	fmt.Println(result)
}

func getPointerValue(a *int) int {
	return *a
}