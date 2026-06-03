package main

import "fmt"

func main() {
	matriz := [2][2]int{
		{1, 2},
		{4, 5},
	}

	for i := range matriz {
		for j := range matriz[i] {
			matriz[i][j] += 2
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}
}
