package main

import "fmt"

func main() {
	//suma
	x := 3000
	y := 4000
	fmt.Println("la suma es: ", x+y)

	//resta
	resultado := x - y
	fmt.Println("El resultadod e la resta es: ", resultado)

	//multiplicacion
	multiplicacion := x * y
	println("El resultado de la multiplicacion es: ", multiplicacion)

	//division
	division := x / y
	modulo := x % y
	fmt.Println("La division es: ", division, " El modulo es: ", modulo)

	//incremento y decremento

	x++
	y++
	fmt.Println(x, y)

	y--
	x--
	fmt.Println(x, y)

}
