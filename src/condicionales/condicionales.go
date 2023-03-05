package main

import (
	"fmt"
	"time"
)

func condicional(numero int) {
	if numero%2 == 0 {
		fmt.Printf("%v es par\n", numero)
	} else {
		fmt.Printf("%v es impar\n", numero)
	}

	if numero%4 == 0 {
		fmt.Printf("%v es divisible por 4\n", numero)
	}
	if numero < 0 {
		fmt.Printf("%v es negativo\n", numero)
	}
	if (numero > 0 && numero < 10) || (numero < 0 && numero > -10) {
		fmt.Printf("%v tiene un dígito\n", numero)
	} else {
		fmt.Printf("%v tiene varios dígitos\n", numero)
	}
}

func swtich_basico() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Hoy es fin de semana")
	default:
		fmt.Println("Hoy es un día laborable")
	}
}

func main() {
	var numero int

	fmt.Print("Ingrese un entero: ")
	fmt.Scanln(&numero)

	condicional(numero)

	swtich_basico()
}
