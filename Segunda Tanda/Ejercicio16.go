package main

import (
	"fmt"
)

func main() {
	var usuarioin string
	var usuariosys string = "acaino"
	var contraseñain string
	var contraseñasys string = "acaino123"

	fmt.Printf("Ingrese nombre de usuario: ")
	fmt.Scan(&usuarioin)

	if usuarioin != usuariosys {
		fmt.Printf("Usuario incorrecto")
	} else {
		fmt.Printf("Ingrese contraseña: ")
		fmt.Scan(&contraseñain)
		if contraseñain != contraseñasys {
			fmt.Printf("Contraseña incorrecta")
		} else {
			fmt.Printf("Acceso otorgado")
		}
	}

	// Buenísimo. Dos comentarios:
	// 1. Los nombres de las variables no son lo mejor pero me parecen correctos solo mantendria el cammelCase.
	// Siendo así las variables deberían haberse llamada usuarioIn, usuarioSys, contraseñaIn, contraseñaSys
	// 2. Dado que usuariosys y contraseñasys son valores fijos contra los que se desea comparar, podrían ser constantes tranquilamente
}
