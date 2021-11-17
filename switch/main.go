package main

import "fmt"

func main() {

	// swith con condicion
	switch modulo := 31 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// switch sin condicion
	value := 120
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 100")
	default:
		fmt.Println("No coinciden")
	}
}
