package main

import "fmt"

func main() {
	var n1 int = 15

	if n1 > 20 {
		fmt.Println("Maior que 20")
	} else if n1 > 15 {
		fmt.Println("Maior que 15")
	} else if n1 > 10 {
		fmt.Println("Maior que 10")
	} else {
		fmt.Println("Número Inválido")
	}
}
