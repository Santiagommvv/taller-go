package main

import (
	"fmt"
	"funciones/funciones"
)

func main() {
	funciones.Saludar("Ana")

	saludo := funciones.Saludar2("Juan")
	fmt.Println(saludo)

	x := "Mundo"
	y := "Hola"
	fmt.Println("x:", x, "| y:", y)
	x, y = funciones.Swap(x, y)
	fmt.Println("Swap => x:", x, "| y:", y)

	resultado, resto := funciones.Dividir(10, 3)
	fmt.Printf("Division: %d. Resto: %d\n", resultado, resto)

	total := funciones.Sumar(10, 10, 10)
	fmt.Println("El total de la suma es:", total)
}
