package main

import (
	"fmt"
)

func main() {
	var numero1 int
	var numero2 int
	var numero3 int

	fmt.Printf("Ingrese un número: ")
	fmt.Scan(&numero1)
	fmt.Printf("Ingrese otro número: ")
	fmt.Scan(&numero2)
	fmt.Printf("Uno más vamos...  : ") // jeje
	fmt.Scan(&numero3)

	if numero1 > numero2 && numero1 > numero3 {
		fmt.Printf("El mayor es %d\n", numero1)
	}
	if numero2 > numero1 && numero2 > numero3 {
		fmt.Printf("El mayor es %d\n", numero2)
	}
	if numero3 > numero1 && numero3 > numero2 {
		fmt.Printf("El mayor es %d\n", numero3)
	}

	// Correcto a medias. Qué pasaría si los tres números son iguales?
}
