package main

import (
	"fmt"
	"math"
)

func main() {

	var numero1, numero2, numero3 float64
	var x1, x2 float64
	var a, b, bPot, c, d, dSqrt float64

	fmt.Printf("Ingrese 3 números reales coeficientes de una función cuadrática a, b, c: ")
	fmt.Scan(&numero1, &numero2, &numero3)

	a = numero1 //coeficientes // Si creés que hace falta especificar que son coeficientes quiere decir que los nombres no son claros o tal vez si pero estás sobre-documentando el código
	b = numero2
	c = numero3

	bPot = math.Pow(b, 2)
	d = bPot - 4*a*c //discriminante // Lo mismo que arriba, solo que en este caso creo que sería mucho más expresivo si en lugar de "d" lo llamaras "discriminante", que te parece?
	dSqrt = math.Sqrt(d)
	x1 = (-b + dSqrt) / (2 * a) //raíz 1  // Idem, sobre-documentación o variables poco expresivas?
	x2 = (-b - dSqrt) / (2 * a) //raíz 2

	if d < 0 {
		fmt.Printf("La función f(x): %.0fx^2 + %.0fx + %.0f no tiene raíces reales", a, b, c)
	}

	if d == 0 {
		fmt.Printf("Las raíces de la función f(x): %.0fx^2 + %.0fx + %.0f son:\nx1=x2=%.0f", a, b, c, x1)
	}

	if d > 0 {
		fmt.Printf("La raíces de la función f(x): %.0fx^2 + %.0fx + %.0f son:\nx1 = %.0f\nx2 = %.0f", a, b, c, x1, x2)
	}

	if a == 0 {
		fmt.Printf("La función f(x): %.0fx + %.0f no es una función cuadrática", b, c)
	}
	//(-b + (b ^ 2 - 4*a*c) ^ (1 / 2)) / (2 * a)

	// Como norma general, si las condiciones son excluyentes unas de otras, osea si solo debería entrar en uno de los ifs, hay que encadenarlas usando un
	// else if. Esto no solo evita que por accidente pueda ingresar en dos condiciones, sino que además vuelva más claro el código. El que lee sabe que solo
	// una condición puede darse y no más.
	// Por el otro lado hay un error en la lógica ya que calculás el discriminante antes de saber si "a" es 0 o no y eso puede traerte problemas. Solo deberías hacer
	// las operaciones que tengan sentido hacer
}
