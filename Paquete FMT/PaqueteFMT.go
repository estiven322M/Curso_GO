package main

import "fmt"

func main() {
	//declaracion de variables
	helloMensaje := "Hello"
	worldMensaje := "Mundo"
	edad := 29

	fmt.Println(helloMensaje, " ,", worldMensaje)

	//printf
	nombre := "Jhon"
	profesion := "Ingeniero de Sistemas"

	fmt.Printf("%s es de profesion %s\n", nombre, profesion)

	//Sprintf
	mensaje := fmt.Sprintf("%s es de profesion %s\n", nombre, profesion) //aqui guarda el string en una variable. Usando Sprintf
	fmt.Println(mensaje)

	//Tipo de dato
	fmt.Printf("Hola Jhon Stivenson: %T", nombre)

	fmt.Printf("Edad : %T", edad)

}
