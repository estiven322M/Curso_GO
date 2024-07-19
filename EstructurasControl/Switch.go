package main

import "fmt"

func main() {
	modulo := 5 % 2

	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No hay condicion")
	}

	//uso de Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y break
	for i := 0; i < 31; i++ {

		//continue
		if i%2 == 0 {
			fmt.Println("Es par")
			continue
		} else {
			fmt.Println("Es impar ", i)
		}
	}
}
