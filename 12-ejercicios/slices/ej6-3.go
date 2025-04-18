package slices

import (
	"fmt"
	pkgslices "slices"
)

func EjS() {
	//invertir, rotar, eliminar f pos, eliminar duplicados
	slice1 := []int{1, 2, 4, 8, 16, 32, 64, 128, 256}
	fmt.Print(slice1, "\ntras invertir se ve como\n")
	fmt.Print(Invertir(slice1), "\n\n")

	fmt.Print(slice1, "\nrotando 13 posiciones\n")
	fmt.Print(Rotar(slice1, 13), "\n\n")

	fmt.Print(slice1, "\ntras borrar cuarto y sexto elementos\n")
	slice1 = Eliminar(slice1, 3)
	slice1 = Eliminar(slice1, 4)
	fmt.Print(slice1, "\n\n") // 0 1 2 -- 3 4 (sexto al momento del print)

	slice2 := []int{1, 64, 4, 64, 16, 32, 64, 128, -64, 1, 1, 32, 48, 100, -64, -63}
	fmt.Print("le borro los duplicados a\n", slice2, "\ny queda\n")
	slice2 = EliminarDuplicados(slice2)
	fmt.Print(slice2, "\n\n")
}

func Invertir(slice []int) []int {
	invertido := make([]int, len(slice), cap(slice))
	for i, v := range slice {
		invertido[len(slice)-i-1] = v
	}
	return invertido
}

func Rotar(slice []int, posiciones int) []int {
	len := len(slice)
	copia := make([]int, len)

	if posiciones < 0 {
		copia = rotarIzq(slice, posiciones)
	} else {
		for i, val := range slice {
			copia[(i+posiciones)%len] = val
		}
	}
	return copia
}

func rotarIzq(slice []int, posiciones int) []int {
	return Invertir(Rotar(Invertir(slice), posiciones*-1))
}

func Eliminar(slice []int, pos int) []int {

	s1 := slice[0:pos]
	s2 := slice[pos+1:]
	copia := pkgslices.Concat(s1, s2)
	//copia = append(copia, 1) // EH?? y esto de dónde lo saqué?

	return copia
}

func EliminarDuplicados(slice []int) []int {
	mapaComparador := make(map[int]bool, len(slice))
	sliceCompactado := []int{}

	for _, v := range slice {
		if !mapaComparador[v] {
			mapaComparador[v] = true
			sliceCompactado = append(sliceCompactado, v)
		}
	}

	return sliceCompactado
}
