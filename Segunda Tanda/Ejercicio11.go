package main

import (
	"fmt"
)

func main() {
	var numero, funcion int

	fmt.Printf("Ingrese un número: ")
	fmt.Scan(&numero)

	funcion = numero*numero + 5*numero - 4 // fijate como el formatter te junta las operaciones según su precedencia

	fmt.Printf("El resultado es: %d", funcion)
}
