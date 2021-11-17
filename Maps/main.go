package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["jose"] = 14
	m["angel"] = 20

	fmt.Println(m)

	// Recorer el diccionario
	for i, v := range m {
		fmt.Println(i, v)
	}

	// validar y acceder mediante la llave del diccionario
	// variable ok nos dice si existe o no existe la key dentro del diccionario (se le puede cambir el nombre al diccionario)
	value, ok := m["angels"]
	fmt.Println(value, ok)
}
