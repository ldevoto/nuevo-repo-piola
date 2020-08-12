package main

import (
	"fmt"
)

func main() {
	var number1, number2, number3, number4 int // No me hagas trampa eh. Cada declaración en una linea nueva =D

	fmt.Printf("Ingrese cuatro números: ")
	fmt.Scan(&number1, &number2, &number3, &number4)

	var first int = number1
	var second int = number2
	var third int = number3
	var fourth int = number4
	var aux int

	if fourth > third {
		aux = third
		third = fourth
		fourth = aux
	}

	if third > second {
		aux = second
		second = third
		third = aux
	}

	if second > first {
		aux = first
		first = second
		second = aux
	}

	if fourth > third {
		aux = third
		third = fourth
		fourth = aux
	}

	if third > second {
		aux = second
		second = third
		third = aux
	}

	if fourth > third {
		aux = third
		third = fourth
		fourth = aux
	}

	fmt.Printf("El orden es: %d - %d - %d - %d", first, second, third, fourth)

	// Me encantó el algoritmo y lo voy a subir como solución sugerida
}
