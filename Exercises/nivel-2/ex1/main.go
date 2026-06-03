package main

import (
	"fmt"
)

func main() {
	// Número de exemplo
	numero := 42

	// %d: Decimal (Base 10)
	// %b: Binário (Base 2)
	// %x: Hexadecimal (Base 16, minúsculo)
	// %X: Hexadecimal (Base 16, maiúsculo)
	// %#x: Hexadecimal com prefixo 0x

	fmt.Printf("Decimal:     %d\n", numero)
	fmt.Printf("Binário:     %b\n", numero)
	fmt.Printf("Hex (min):   %x\n", numero)
	fmt.Printf("Hex (mai):   %X\n", numero)
	fmt.Printf("Hex (c/pre): %#x\n", numero)
}
