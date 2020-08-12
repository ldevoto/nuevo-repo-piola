package main

import (
	"fmt"
)

func main() {
	var numero int

	fmt.Printf("Ingrese un número: ")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Printf("El número es par")
	} else {
		fmt.Printf("El número es impar")
	}
}
