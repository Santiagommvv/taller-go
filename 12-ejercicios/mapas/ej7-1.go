package mapas

import (
	"fmt"
	"strings"
)

func EjM() {
	//frecuencia de palabras
	refranCinco := "son las cinco menos cinco y faltan cinco para decir son las cinco entonces cuantas son las veces que dije cinco sin contar el ultimo cinco"
	fmt.Print(refranCinco, "\n")
	fmt.Print("frecuencia de palabras\n", ContarPalabras(refranCinco), "\n")

	//mapas iguales?

	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6}
	m2 := map[string]int{"d": 4, "e": 5, "b": 2, "a": 1, "f": 6, "c": 3}
	fmt.Print("\n[a:1 b:2 c:3 d:4 e:5 f:6]")
	fmt.Print("\n[d:4 e:5 b:2 a:1 f:6 c:3]\n")
	fmt.Printf("Son iguales : %v\n\n", Igual(m1, m2))

	//anagramas?
	sBase := "animal inala m"
	sVar1 := "lami na aamiln"
	fmt.Print("string1: ", sBase, "\nstring2: ", sVar1, "\n")
	fmt.Print("Son anagramas : ", Anagramas(sVar1, sBase), "\n")
}

func ContarPalabras(str string) (mapa map[string]int) {
	palabras := strings.Split(str, " ")
	mapa = make(map[string]int)

	for _, elemento := range palabras {
		if mapa[elemento] != 0 {
			mapa[elemento]++
		} else {
			mapa[elemento] = 1
		}
	}

	return mapa
}

// mapas de mismas claves y valores
func Igual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for clave, valor := range x {
		valorEnY, ok := y[clave]
		if !ok || valorEnY != valor {
			return false
		}
	}
	return true
}

// mapas equivalentes con clave caracter
func igualCaracteres(x, y map[rune]int) bool {
	if len(x) != len(y) {
		return false
	}

	for clave, valorX := range x {
		valorEnY, ok := y[clave]
		if !ok || valorEnY != valorX {
			return false
		}
	}
	return true
}

func Anagramas(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	mapa1 := make(map[rune]int) //mapa de frecuencia de letras, no de palabras
	mapa2 := make(map[rune]int)

	for _, caracter := range s1 {
		mapa1[caracter]++
	}

	for _, caracter := range s2 {
		mapa2[caracter]++
	}

	return igualCaracteres(mapa1, mapa2)
}
