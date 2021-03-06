package main

import (
	"fmt"
)

func main() {
	// Declarar constantes

	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declarar variables

	age := 18 // Dos puntos se usan cuando la variablo no esta creada entonces se creara y se le asignara el valor y parsea el tipo de dato
	var area int = 10
	var area2 = 10

	fmt.Println(age, area, area2)

	// Zero values
	var a int     // enteras retorna 0
	var b float64 // decimales retorna 0
	var c string  // texto retorna empty
	var d bool    // true o false retorna false

	fmt.Println(a, b, c, d)

	// area de un cuadrado
	const base float64 = 10.5
	const altura float64 = 10.5

	areaCuadrado := base * altura

	fmt.Println("area:", areaCuadrado)

}
