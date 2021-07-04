package main

import "fmt"

func main() {

	var x [5]int
	x[0] = 1
	x[1] = 2
	x[2] = 3
	fmt.Println(x)

	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y[1])

}
