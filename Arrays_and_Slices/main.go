package main

import "fmt"

func main() {

	// Definir un Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// slice
	slice := []int{0, 0, 1, 3, 5, 3}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos slice
	// slice[primer indice: segundo indice]
	// primer indice: es inclusivo, es decir que si se cuenta al tomar el rango de valors
	// segundo indice: es exclusivo, es decir que no se cuenta al tomar el rango de valores
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:5])
	fmt.Println(slice[0:])

	// Append
	slice = append(slice, 10)
	fmt.Println(slice)

	// Append new list
	newList := []int{8, 9, 10}
	slice = append(slice, newList...)
	fmt.Println(slice)
}
