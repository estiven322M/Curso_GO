package main

import "fmt"

func numerosParesImpares(arreglo []int) {

	var contNumerosPares uint
	var contNumerosImpares uint

	numerosPares := []int{}
	numerosImpares := []int{}

	for i := 0; i < len(arreglo); i++ {

		if arreglo[i]%2 == 0 { //validamos si el numero es par
			contNumerosPares++
			numerosPares = append(numerosPares, arreglo[i])
		} else {
			contNumerosImpares++
			numerosImpares = append(numerosImpares, arreglo[i])
		}
	}
	fmt.Println("El slicen de numeros enteros pares es: ", numerosPares)
	fmt.Println("El slicen de numeros enteros impares es: ", numerosImpares)
	fmt.Println("El slicen original tiene un total de :", len(arreglo), " elementos y un total de ", contNumerosPares, " numeros pares y ", contNumerosImpares, " numeros impares")
}
