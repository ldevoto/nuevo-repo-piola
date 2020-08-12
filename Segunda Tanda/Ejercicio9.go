package main

import (
	"fmt"
)

func main() {
	var NumeroSemana int

	fmt.Printf("Ingrese un número: ")
	fmt.Scan(&NumeroSemana)

	switch NumeroSemana {
	case 0:
		fmt.Printf("domingo")
	case 1:
		fmt.Printf("lunes")
	case 2:
		fmt.Printf("martes")
	case 3:
		fmt.Printf("miércoles")
	case 4:
		fmt.Printf("jueves")
	case 5:
		fmt.Printf("viernes")
	case 6:
		fmt.Printf("sábado")
	default:
		fmt.Printf("error")
	}

	// Joya nunca taxi nunca remis
}
