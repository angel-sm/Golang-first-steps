package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {

	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// otra forma de declarar una clase
	var otherCar car
	otherCar.brand = "Chevrolet"
	fmt.Println(otherCar)
}
