package main

import "fmt"

func main() {

	//calculo area de un triangulo
	var alturaTriangulo int = 34
	var baseTriangulo int = 18
	var areaTriangulo float64 = float64(alturaTriangulo) * float64(baseTriangulo) * 0.5
	fmt.Println("El area del triangulo es: ", areaTriangulo)

	//calculo de la area de un trapecio
	var base1 float64 = 34
	var base2 float64 = 67
	var altura float64 = 12

	area := 0.5 * (base1 + base2) * altura

	fmt.Println("El area del trapecio es: ", area)

}
