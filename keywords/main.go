package main

import "fmt"

func main() {

	// Defer: esta sentencia se ejecutara al final de todo el codigo
	// Se puede usar para cuando se llama a la db al final cerrar la conexion
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue se usa para que el siclo aun teniendo un error, pero que sea error controlado
		// continue el siclo
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// Break se usa para romper el siclo, en este caso cuando la condicion se cumpla
		if i == 7 {
			fmt.Println("Break")
			break
		}
	}
}
