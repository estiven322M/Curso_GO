package main

import "fmt"

func main() {
	var dog, cat = "Waw Waw Waw soy un perro muy bravo", "Miau Miau soy un gato muy perezoso"

	fmt.Println(dog, cat) //imprimir el valor

	const pi float64 = 3.14

	fmt.Println("pi: ", pi) //imprimir el valor

	var altura int = 14
	fmt.Println(altura)

	//Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio: calculo de área de un cuadrado

	const baseCuadrado0 = 10
	areaCuadrado := baseCuadrado0 * baseCuadrado0
	fmt.Println("El area del cuadrado es: ", areaCuadrado)

}
