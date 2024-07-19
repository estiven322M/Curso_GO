package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"Hola", "Jhon", "Stivenson", "Mendez", "Alvarez"}

	for i, valor := range slice {
		fmt.Println(i, valor)

	}

	isPalindromo("Ama")

}

// funcion palindromo
func isPalindromo(text string) {

	minuscula := strings.ToLower(text)
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(minuscula[i])
	}

	if minuscula == textReverse {
		fmt.Println("Es palindromo")

	} else {
		fmt.Println("No es palindromo")
	}
}
