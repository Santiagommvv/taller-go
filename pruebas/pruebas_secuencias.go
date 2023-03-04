package main

import (
	"fmt"
	"secuencias"
)

func main() {
	arreglo := []int{-1, 5, -6, 10, 7, 2, -10, 5, 3, 4}
	sumaMaxima, posInicial, posFinal := secuencias.SubsecuenciaSumaMaxima(arreglo)
	fmt.Printf("La suma máxima es: %d\n", sumaMaxima)
	fmt.Printf("Posición Inicial: %d\n", posInicial)
	fmt.Printf("Posición Final: %d\n", posFinal)
}
