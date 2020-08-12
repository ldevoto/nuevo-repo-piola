package main

import (
	"fmt"
)

func main() {
	var mesdelaño string

	fmt.Printf("Ingrese el mes deseado: ")
	fmt.Scan(&mesdelaño)

	switch mesdelaño {
	case "Enero":
		fmt.Printf("1")
	case "Febrero":
		fmt.Printf("2")
	case "Marzo":
		fmt.Printf("3")
	case "Abril":
		fmt.Printf("4")
	case "Mayo":
		fmt.Printf("5")
	case "Junio":
		fmt.Printf("6")
	case "Julio":
		fmt.Printf("7")
	case "Agosto":
		fmt.Printf("8")
	case "Septiembre":
		fmt.Printf("9")
	case "Octubre":
		fmt.Printf("10")
	case "Noviembre":
		fmt.Printf("11")
	case "Diciembre":
		fmt.Printf("12")
	default:
		fmt.Printf("Error")
	}

	// Espléndido
}
