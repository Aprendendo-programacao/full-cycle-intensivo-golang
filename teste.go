package main

import "fmt"

func main() {

	mapExample := make(map[string]int)

	mapExample["wesley"] = 10
	mapExample["mariana"] = 12

	fmt.Println(mapExample["wesley"])

}
