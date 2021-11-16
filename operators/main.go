package main

import (
	"fmt"
	"math"
)

func main() {
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
