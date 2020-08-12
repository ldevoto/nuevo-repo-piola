package main

import (
	"fmt"
)

func main() {
	var lado1, lado2, lado3 float64

	fmt.Printf("Ingrese 3 lados de un triángulo: ")
	fmt.Scan(&lado1, &lado2, &lado3)

	var a float64 = lado1
	var b float64 = lado2
	var c float64 = lado3

	if a == b && b == c {
		fmt.Printf("Equilátero")
	} else if a != b && a != c && b != c {
		fmt.Printf("Escaleno")
	} else {
		fmt.Printf("Isósceles")
	}

	// Beautiful. Solo no entiendo esa asignación de lado1, lado2 y lado3 a a, b y c respectivamente
}
