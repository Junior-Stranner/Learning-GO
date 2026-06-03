package main

import (
	"fmt"
)

func main() {
	var numero1, numero2 int
	numero1 = 10
	numero2 = 5

	fmt.Println("numero1 =", numero1)
	fmt.Println("numero2 =", numero2)

	// Operadores relacionais
	fmt.Println("numero1 == numero2:", numero1 == numero2) // igual
	fmt.Println("numero1 != numero2:", numero1 != numero2) // diferente
	fmt.Println("numero1 > numero2:", numero1 > numero2)   // maior que
	fmt.Println("numero1 < numero2:", numero1 < numero2)   // menor que
	fmt.Println("numero1 >= numero2:", numero1 >= numero2) // maior ou igual
	fmt.Println("numero1 <= numero2:", numero1 <= numero2) // menor ou igual
}
