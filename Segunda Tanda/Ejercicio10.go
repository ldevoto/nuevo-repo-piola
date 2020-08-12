package main

import (
	"fmt"
)

func main() {
	var numero, funcion int

	fmt.Printf("Ingrese un número: ")
	fmt.Scan(&numero)

	funcion = numero*2 + 10

	fmt.Printf("El resultado es: %d", funcion)

	// easy peacy no?
}
