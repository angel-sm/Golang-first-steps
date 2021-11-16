package main

import "fmt"

func main() {

	// PrintLn: agrega un salto de linea al final de cada impresion
	helloMessage := "Hola"
	wordMessage := "word"
	fmt.Println(helloMessage, wordMessage)
	fmt.Println(helloMessage, wordMessage)

	// Printf para agregar formato de datos
	name := "Jose"
	age := 19

	fmt.Printf("%s tienen %d de edad\n", name, age)
	// Cuando no sabemos el tipo de dato que ira se agrega %v
	fmt.Printf("%v tienen %v de edad\n", name, age)

	// Conocer el tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("age: %T\n", age)

}
