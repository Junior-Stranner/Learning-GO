package main

import "fmt"

func main() {
	var n1, n2 int
	n1 = 10
	n2 = 5

	//Operadores Lógicos = (&& - and || - or ! - not)

	//Operadores AND - && (v AND v) = true
	fmt.Println(n1 != n2 && n1 > n2)

	//Operadores OR - || (v OR v) = true
	fmt.Println(n1 == n2 && n1 > n2)

	//Operadores NOT - || (v NOT v) = Inverso
	fmt.Println(!(n1 == n2))

}
