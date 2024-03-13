package ciclos

import (
	"fmt"
)

func Java() {
	// A la Java
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\t", i)
	}
	fmt.Println()
}

func While() {
	//Como ciclo while
	i := 0
	for i < 10 {
		fmt.Printf("%v\t", i)
		i++
	}
	fmt.Println()
}
func Infinito() {
	// Bucle infinito
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Printf("%v\t", i)
		i++
	}
	fmt.Println()
}

func Range_arreglo() {
	// Range con arreglos
	arreglo := [10]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	for indice, valor := range arreglo {
		fmt.Printf("2**%v = %v\n", indice, valor)
	}
}

func Range_string() {
	// Range con strings
	cadena := "Algoritmos y Programación II"
	for _, letra := range cadena {
		fmt.Printf("%c\t", letra)
	}
}
