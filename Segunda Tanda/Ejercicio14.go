package main

import (
	"fmt"
	"math"
)

func main() {
	var radio, radiocuadrado, radiocubo float64

	fmt.Printf("Ingrese el radio de la circunferencia: ")
	fmt.Scan(&radio)

	var circ, sup, vol float64

	circ = math.Pi * 2 * radio

	radiocuadrado = math.Pow(radio, 2)
	sup = math.Pi * radiocuadrado // esta superficie no es la de un círculo? me mareé jajja

	radiocubo = math.Pow(radio, 3)
	vol = (4 / 3) * math.Pi * radiocubo

	if radio == 0 {
		fmt.Printf("Con radio 0, todo es 0")
	} else if radio < 0 {
		fmt.Printf("Ya nada tiene sentido")
	} else {
		fmt.Printf("La circunferencia es %.2f, la superficie es %.2f, el volumen es %.2f", circ, sup, vol) //no se si querés que se exprese en unidades
	}

	// Al igual que antes, si el radio es 0 o menor que 0 no hace falta que hagas ningún cálculo, solo deberías hacerlos en caso de que radio > 0
}
