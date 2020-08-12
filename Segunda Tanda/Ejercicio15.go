package main

import (
	"fmt"
)

func main() {
	var numero1, numero2, numero3, numero4, numero5, division1, division2, division3, division4, division5 float64

	fmt.Printf("Ingrese 5 números reales: ")
	fmt.Scan(&numero1, &numero2, &numero3, &numero4, &numero5)

	if numero2 == 0 || numero3 == 0 || numero4 == 0 || numero5 == 0 {
		fmt.Printf("No voy a caer...")
	} else {
		division1 = numero1 / 1
		division2 = division1 / numero2
		division3 = division2 / numero3
		division4 = division3 / numero4
		division5 = division4 / numero5
		fmt.Printf("El resultado es %f", division5)
	}

	// Joyaa
}
