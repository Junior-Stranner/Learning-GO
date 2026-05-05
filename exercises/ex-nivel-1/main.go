package main

import "fmt"

func main() {
	// Variáveis
	var nome string = "Go"
	idade := 15 // inferência de tipo

	fmt.Println("Linguagem:", nome)
	fmt.Println("Anos desde criação:", idade)

	// Constantes
	const pi = 3.14
	fmt.Println("Pi:", pi)

	// Tipos básicos
	var inteiro int = 42
	var decimal float64 = 3.14
	var booleano bool = true
	var texto string = "Olá, mundo!"

	fmt.Println(inteiro, decimal, booleano, texto)
}
