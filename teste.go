package main

import (
	"fmt"
)

type vehicle interface {
	start() string
}

type Car struct {
	CarName string `json:"car"`
	CarYear int    `json:"year"`
}

func (car Car) start() string {
	return "Iniciou"
}

func (car Car) drive() {
	fmt.Println(car.CarName, "andou!")
}

func acceptOnlyVehicleImplementation(car vehicle) {
	fmt.Println("É um veículo")
}

func main() {

	car := Car{
		CarName: "Fusion",
		CarYear: 2020,
	}

	car.drive()
	acceptOnlyVehicleImplementation(car)

}
