package main

import "fmt"

type CarName string
type CarYear int

func main() {

	var carName CarName
	carName = "Fusion"
	fmt.Println(carName)

	var carYear CarYear
	carYear = 2020
	fmt.Println(carYear)

}
