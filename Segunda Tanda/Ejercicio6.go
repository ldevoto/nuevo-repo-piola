package main

import (
	"fmt"
)

func main() {
	var numero1, numero2, numero3, numero4 int

	fmt.Printf("Ingresa los 4 números: ")
	fmt.Scan(numero1, numero2, numero3, numero4)

	var Primero int = numero1
	var Segundo int = numero2
	var Tercero int = numero3
	var Cuarto int = numero4

	if numero1 > numero2 && numero1 > numero3 && numero1 > numero4 {
		Primero = numero1 //redundante I know
		if numero2 > numero3 && numero2 > numero4
			Segundo = numero2
			if numero4 > numero3
			Tercero = numero4
			Cuarto = numero3
		else if numero3 > numero2 && numero3 > numero4
			Segundo = numero3
			if numero4 > numero2
			Tercero = numero4
			Cuarto = numero2
		else if numero4 > numero2 && numero4 > numero3
			Segundo = numero4
			if numero
			
	if numero2 > numero1 && numero2 > numero3 && numero2 > numero4 {
		Primero = numero2
		if 
	}
	if numero3 > numero1 && numero3 > numero2 && numero3 > numero4 {
		Primero = numero3
	}
	if numero4 > numero1 && numero4 > numero2 && numero4 > numero3 {
		Primero = numero4
		
	// Bueno este puedo asumir que quedó colgado. Tiene problemas de compilación y de lógica interesantes jajaja
}
