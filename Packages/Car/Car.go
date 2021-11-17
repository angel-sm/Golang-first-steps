package car

import "fmt"

// Para hacer un struc privado se inicia con una minuscula y cuando es publico se incializa con una mayuscula
// al igual que sus atributos
// al igual que sus metodos

// CarPublic con acceso publico
type CarPublic struct {
	// Para que las variables sean publicas deben de ser en mayusculas
	Brand string
	Year  int
}

type carPrivate struct {
	// Para que las variables sean publicas deben de ser en mayusculas
	brand string
	year  int
}

func printMessage(message string) {
	fmt.Println(message)
}

func PrintMessage(message string) {
	fmt.Println(message)
}
