package main

import (
	"fmt"
)

func main() {
	var numero1 int
	var numero2 int

	fmt.Printf("Ingrese un número: ")
	fmt.Scan(&numero1)
	fmt.Printf("Ingrese otro número: ")
	fmt.Scan(&numero2)

	if numero1 > numero2 {
		fmt.Printf("El mayor es %d", numero1)
	} else if numero2 > numero1 {
		fmt.Printf("El mayor es %d", numero2)
	}

	// apa, que ordenadito papáaa. Me gusta :+1:
}
