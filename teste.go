package main

import "fmt"

func main() {

	slice := make([]int, 5)
	fmt.Println(slice)

	slice[0] = 1
	slice[1] = 1
	slice[2] = 1
	slice[3] = 1
	slice[4] = 1

	slice = append(slice, 1, 2, 3, 4)
	fmt.Println(slice)

	sliceString := []string {
		"Hello",
		"World",
	}
	fmt.Println(sliceString[0])

}
