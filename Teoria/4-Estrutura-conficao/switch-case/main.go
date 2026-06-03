package main

import "fmt"

func main() {
	var n1 int = 15

	switch {
	case n1 > 20:
		fmt.Println("Maior que 20")
	case n1 > 15:
		fmt.Println("Maior ou igual a 15")
	case n1 > 10:
		fmt.Println("Maior que 10")
	default:
		fmt.Println("Número inválido")
	}
}
