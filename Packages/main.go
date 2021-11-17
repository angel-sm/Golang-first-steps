package main

import (
	"fmt"
	car "practicas_go/Packages/Car"
)

func main() {
	var myCar car.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	car.PrintMessage("Salu2 metodo publico")
}
