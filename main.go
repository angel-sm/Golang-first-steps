package main

import (
	"fmt"
	"math"
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

	// Operadores aritmeticos

	x := 10
	y := 5

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = x - y
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Divicion
	result = x / y
	fmt.Println("Divicion:", result)

	// Modulo
	result = x % y
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	const baseRectangle int = 10
	const heigthRectangle int = 20

	// Area rectangulo
	result = baseRectangle * heigthRectangle
	fmt.Println("Area rectangulo:", result)

	// Area trapecio
	const heigthTrapecio int = 4
	const baseBottomTrapecio int = 10
	const baseTopTrapecio int = 4

	result = ((baseBottomTrapecio + baseTopTrapecio) * heigthTrapecio) / 2
	fmt.Println("Area trapecio:", result)

	// Area Circulo
	const radioCirculo = 19
	areaCirculo := math.Pi * math.Pow(radioCirculo, 2)
	fmt.Println("Area Circulo:", areaCirculo)

}
