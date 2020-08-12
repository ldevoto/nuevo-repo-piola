package main

import (
	"fmt"
)

func main() {
	var numero1, numero2, Aux int // Una linea por variables y los nombres siempre en minúsculas

	fmt.Printf("Ingrese dos números: ")
	fmt.Scan(&numero1, &numero2)

	if numero1 < numero2 {
		Aux = numero1
		numero1 = numero2
		numero2 = Aux
	}

	if numero1-numero2 >= 130 {
		fmt.Printf("Están distanciados")
	} else if numero1-numero2 >= 30 {
		fmt.Printf("Están cerca")
	} else if numero1-numero2 >= 1 {
		fmt.Printf("Están re cerquita")
	} else if numero1-numero2 == 0 {
		fmt.Printf("Son iguales")
	}

	// Joya, lo veo bien. Hay dos leves cambios que le haría:
	// 1. Guardaría el valor de la resta en una variables temporal tipo "diferencia" y usaría esa para todas las comparaciones en lugar de calcular siempre lo mismo
	// 2. El último if de todos es redundante, lo sacaría y dejaría solo un else
}
