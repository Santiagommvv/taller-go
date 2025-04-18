package main

import (
	"fmt"

	misMapas "github.com/untref-ayp2/taller-GO/12-ejercicios/mapas"
	misSlices "github.com/untref-ayp2/taller-GO/12-ejercicios/slices"
)

// Llamo todos los ejercicios desde acá en un método expuesto por tema/paquete
func main() {

	fmt.Print("SLICES\n******\n")
	misSlices.EjS()

	fmt.Print("MAPAS\n*****\n")
	misMapas.EjM()
}
