package main

import "fmt"

func main() {

	/*	numeros := [5]int{1, 2, 3, 4, 5}

		for i := 0; i < len(numeros); i++ {
			numeros[i] += 2
			fmt.Println(numeros[i])
		}*/

	/*
	   var numeros []int

	   numeros = append(numeros, 1, 2, 3, 4, 5, 6, 7)
	   fmt.Println(numeros)
	*/
	var numeros []int

	numeros = append(numeros, 1, 2, 3, 4, 5, 6, 7)

	numerosFiltrados := numeros[:6]

	fmt.Println(numerosFiltrados)
	fmt.Println(numeros)

}
