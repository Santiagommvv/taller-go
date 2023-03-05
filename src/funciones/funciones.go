package funciones

import "fmt"

func Saludar(persona string) {
	fmt.Println("Hola", persona)
}

func Saludar2(persona string) string {
	saludo := "Hola " + persona
	return saludo
}

func Swap(x, y string) (string, string) {
	//Multiple valores de retorno
	return y, x
}

func Dividir(dividendo, divisor int) (resultado, resto int) {
	//Valores de retorno nombrados
	resultado = dividendo / divisor
	resto = dividendo % divisor
	return
}
