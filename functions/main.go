package main

import "fmt"

func print(message string) {
	fmt.Println(message)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnDobleArgument(a int) (b, c int) {
	return a, a * 2
}

func main() {
	print("Hola")
	print("Mundo")
	tripeArgument(1, 2, "Hola")

	getFirtsArgument, _ := returnDobleArgument(5)
	fmt.Println("firts Argument:", getFirtsArgument)

	_, getSecondArgument := returnDobleArgument(5)
	fmt.Println("second Argument:", getSecondArgument)

	getFirtsArgument, getSecondArgument = returnDobleArgument(5)
	fmt.Println("first and second Argument:", getFirtsArgument, getSecondArgument)
}
