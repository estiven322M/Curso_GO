package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	//Iniciamos el Random
	rand.Seed(time.Now().UnixNano())

	//generar un numero aleatorio

	contador := 0
	valor := rand.Intn(100)
	encontrado := false

	for contador != valor {
		contador = rand.Intn(100)
		if contador == valor {
			encontrado = true

			break
		}
		contador++

	}
	if encontrado {
		fmt.Printf("Encontrado. Contador es: ", contador, valor)
	} else {
		fmt.Println("No se encontro el valor en 100 intentos")
	}

	//Convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)

}
