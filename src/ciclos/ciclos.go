package main

import "fmt"

func java() {
	// A la Java
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\t", i)
	}
}

func while() {
	//Como ciclo while
	i := 0
	for i < 10 {
		fmt.Printf("%v\t", i)
		i++
	}
}
func infinito() {
	// Bucle infinito
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Printf("%v\t", i)
		i++
	}
}

func range_arreglo() {
	// Range con arreglos
	arreglo := [10]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	for indice, valor := range arreglo {
		fmt.Printf("2**%v = %v\n", indice, valor)
	}
}

func range_string() {
	// Range con strings
	cadena := "Algoritmos y Programación II"
	for _, letra := range cadena {
		fmt.Printf("%c\t", letra)
	}
}
func main() {
	java()
	//while()
	//infinito()
	//range_arreglo()
	//range_string()
}
