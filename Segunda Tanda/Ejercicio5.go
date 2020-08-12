package main

import (
	"fmt"
)

func main() {
	var numero1 int
	var numero2 int
	var numero3 int
	var numero4 int

	fmt.Printf("Ingrese cuatro números intercalados con un guión: \n") // El mensaje quedó mal
	fmt.Scan(&numero1, &numero2, &numero3, &numero4)

	var Máximo int = numero1 // Guarda con los nombres de las variables. Deberían ser siempre en minúsculas y nunca llevar acento

	if numero2 > numero1 && numero2 > numero3 && numero2 > numero4 {
		Máximo = numero2
	}
	if numero3 > numero1 && numero3 > numero2 && numero3 > numero4 {
		Máximo = numero3
	}
	if numero4 > numero1 && numero4 > numero2 && numero4 > numero3 {
		Máximo = numero4
	}

	fmt.Printf("El valor máximo es: %d", Máximo)

	// Ok. Usaste una parte de un tipo de solución mezclada con otra jaja. Arranca definiendo un máximo y asumiendo que el numero1 es el máximo.
	// Y después te olvidás de ese máximo para las comparaciones y lo hacés por fuerza bruta digamos.
	// Lo que hubiese estado mejor sería comparar cada uno de los valores contra el máximo, mirá:

	Máximo = numero1

	if numero2 > Máximo {
		Máximo = numero2
	}
	if numero3 > Máximo {
		Máximo = numero3
	}
	if numero4 > Máximo {
		Máximo = numero4
	}

	fmt.Printf("El valor máximo es: %d\n", Máximo)
}
