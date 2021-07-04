package main

import "fmt"

type Car struct {
	CarName string
	CarYear int
}

func (car Car) drive() {
	fmt.Println(car.CarName, "andou!")
}

func (car *Car) changeCarName()  {
	fmt.Println(car.CarName)
	car.CarName = "Fusion2"
	fmt.Println(car.CarName)
}

func main() {

	car := Car{
		CarName: "Fusion",
		CarYear: 2020,
	}

	car.drive()
	car.changeCarName()

}
