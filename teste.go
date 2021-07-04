package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	CarName string `json:"car"`
	CarYear int    `json:"year"`
}

func (car Car) drive() {
	fmt.Println(car.CarName, "andou!")
}

func main() {

	carJson := []byte(`{"car": "BMW", "year": 2020}`)
	var car1 Car
	json.Unmarshal(carJson, &car1)
	fmt.Println(car1)

	car2 := Car{
		CarName: "Fusion",
		CarYear: 2020,
	}

	car2.drive()

	result, _ := json.Marshal(car2)
	fmt.Println(string(result))

}
