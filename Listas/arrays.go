package main

import "fmt"

func main() {
	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//slice
	slice := []int{0, 2, 4, 6, 8, 9, 12}
	fmt.Println(slice, len(slice), cap(slice))

	//metodos en el slice
	fmt.Println(slice[0])   //imprime la primera posicion
	fmt.Println(slice[:3])  //imprime desde la posicion 0 hasta la posicion 3
	fmt.Println(slice[2:4]) //imprime desde la posicion 2 hasta la posicion 4
	fmt.Println(slice[4:])  //imprime desde la posicion 4 en adelante

	//Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//append nueva lista (agregar una lista)
	newSlice := []int{8, 9, 19}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
