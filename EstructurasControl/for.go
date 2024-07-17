package main

import "fmt"

func main() {

	//ciclo for condicional
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	//For while
	contador := 0
	for contador <= 10 {
		fmt.Println(contador)
		contador++
	}

	// For forever
	//contadorForever := 0
	//for {
	//fmt.Println(contadorForever)
	//contadorForever++
	//}

	// for invertido

	for j := 900; j > 0; j-- {
		fmt.Println(j)
	}

}
