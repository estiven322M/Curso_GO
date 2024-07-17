package main

import "fmt"

func funcionNomal(mensaje string) { //funcion con un solo argumento
	fmt.Println(mensaje)
}
func funcionConArgumentos(a, b int, c string) { //funcion con varios parametros de entrada
	fmt.Println(a, b, c)
}

func funcionConRetorno(a int) int { // funcion con valores de retorno
	return a * 2
}

func dobleRetorno(a int) (c, d int) {
	return a, a * 2
}

func main() {
	funcionNomal("Hola Jhon Stivenson")
	funcionConArgumentos(14, 04, "Jhon Stivenson")
	value := funcionConRetorno(400)
	fmt.Println("El valor de la multiplicacion es: ", value)

	value1, value2 := dobleRetorno(500)
	fmt.Println("value1 y Value2 son: ", value1, value2)

}
