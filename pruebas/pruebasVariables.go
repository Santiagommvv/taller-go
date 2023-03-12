package main

import (
	"fmt"
	"math"
)

var x = 10           //variable global
const pi = 3.1415926 //constante

func main() {
	var cadena = "cadena de caracteres"
	otra_cadena := "Otra cadena"
	var otra_mas string = "Tercera cadena"

	fmt.Println(cadena)
	fmt.Println(otra_cadena)
	fmt.Println(otra_mas)

	var a, b int = 1, 2 //declaración múltiple
	c, ultima_cadena := 10, "Última cadena"

	fmt.Println(a, b)
	fmt.Println(c, ultima_cadena)

	x = 17 //variable global
	fmt.Println("Variable  global:", x)

	var x = 4 //variable local oculta a la variable global x

	for x := 0; x < 10; x++ { //variable del for oculta a la variable local x
		fmt.Print(x, ", ")
	}
	fmt.Println("\nVariable local:", x)
	fmt.Println(pi)

	//constantes predefinidas
	fmt.Println(math.Pi)

	//Números complejos
	var complejo complex64
	complejo = complex(0, 1)
	cuadrado := complejo * complejo

	fmt.Printf("%v ** 2 = %v\n", complejo, cuadrado)
}
