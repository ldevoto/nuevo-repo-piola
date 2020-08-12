package main

import (
	"fmt"
)

func main() {

	var usuarioin string
	var usuariosys string = "acaino"
	var contraseñain string
	var contraseñasys string = "acaino123"
	var quierecambiar string
	var nuevacontraseña string
	var nuevacontraseñaconf string // Lo mismo que en el anterior, todas las variables que estén compuestas por más de una palabra, cammelCase

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
			fmt.Printf("Acceso otorgado\n")
			fmt.Printf("¿Desea cambiar la contraseña? S/N\n")
			fmt.Scan(&quierecambiar)
			if quierecambiar == "N" {
				fmt.Printf("Su contraseña vencerá en 13 días")
			} else {
				fmt.Printf("Ingrese su nueva contraseña: ")
				fmt.Scan(&nuevacontraseña)
				fmt.Printf("Ingrese nuevamente la contraseña: ")
				fmt.Scan(&nuevacontraseñaconf)
				if nuevacontraseña == nuevacontraseñaconf {
					fmt.Printf("Contraseña cambiada exitosamente")
				} else {
					fmt.Printf("Error al cambiar la contraseña. Las contraseñas no coinciden")
				}
			}
		}
	}

	// Joya, nada que modificar.
	// Solo mencionar algo curioso: cuando se le pregunta al usuario si desea cambiar la contraseña y si por algún error el usuario ingresa una "j"
	// se va cambiar. Por lo general se encara al revés. A menos que explicitamente ingreses que querés realizar la acción, nada se hace. Osea, la pregunta
	// debería si quierecambiar == "S" -> ejecuto la acción
}
