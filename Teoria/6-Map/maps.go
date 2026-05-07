package main

import "fmt"

func main() {

	/*	notas := make(map[string]float64)

		//atribuindo valores
		notas["Junior"] = 8.5
		notas["João"] = 9.0

		fmt.Println(notas)

		//verificar se a posicao existe dentro do map
		value, existe := notas["Marco"]

		//remove item do map
		delete(notas, "Marco")

		fmt.Println(value, existe)
	*/

	usuarios := map[string]string{
		"nome":      "Heinz Jr",
		"profissao": "Dev",
	}

	fmt.Println(usuarios)
	fmt.Println(usuarios["nome"])

	usuario2 := map[string]map[string]string{
		"nome": map[string]string{
			"primeiro": "Heinz Jr",
			"ultimo":   "Stranner",
		},
		"profissao": map[string]string{
			"primeiro": "dev",
			"ultimo":   "Junior",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"])
	fmt.Println(usuario2["profissao"]["ultimo"])

}
