package funciones

func Sumar(numeros ...int) (total int) {
	//Función con una cantidad indefinida de parámetros
	total = 0
	for _, valor := range numeros {
		total += valor
	}
	return
}
