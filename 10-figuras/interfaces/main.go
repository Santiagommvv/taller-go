package main

import (
	"fmt"

	"github.com/untref-ayp2/taller-GO/10-figuras/figuras"
)

func mostrar(f figuras.Figura) {
	fmt.Println(f.ToString())
	fmt.Println("Area: ", f.Area())
	fmt.Println("Perimetro: ", f.Perimetro())
}

func main() {
	p1 := figuras.Punto{X: 0, Y: 0}
	p2 := figuras.Punto{X: 10, Y: 5}

	r := figuras.Rectangulo{P1: p1, P2: p2}
	c := figuras.Cuadrado{Pto: p1, Lado: 10}

	arreglo_figuras := [2]figuras.Figura{r, c}

	for _, f := range arreglo_figuras {
		mostrar(f)
	}
}
