package main

import (
	"fmt"
	"strings"
)

func saludo(primerNombre, segundoNombre string) {
	fmt.Println(primerNombre, segundoNombre)
}
func sumas(num1, num2, num3 int) int {
	valor := num1 + num2 + num3
	return valor
}

func retornoVariosValores(texto string) (string, string) {
	minuscula := strings.ToLower(texto)
	mayuscula := strings.ToUpper(texto)
	return minuscula, mayuscula
}

func main() {
	saludo("Jhon", "Stivenson")
	fmt.Println(sumas(2300, 1345, 876))
	fmt.Println(retornoVariosValores("Jhon"))

}
