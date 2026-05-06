package main

import (
	"fmt"
)

func main() {
	//atribuicao simples
	var numero1 int = 10

	//declaração curta
	numero2 := 20

	//+=, -=, *=, /=
	numero2 += 5 //numero2 e somar com 5
	numero2 -= 3 //numero2 e Subtrair com 3
	numero2 *= 2 //numero2 e Multiplicar com 2
	numero2 /= 1 //numero2 e Dividr com 1
	fmt.Println("Divisão = ", float64(numero1)/float64(numero2))

}
