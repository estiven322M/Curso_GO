package main

//proyecto en GO que suma dos numeros. Calculadora

import ( //importacion de librerias para el proyecto
	"bufio" //libreria de entrada por teclado
	"fmt"   //libreria para mostrar texto en la consola
	"math"  //libreria matematica
	"os"
	"strconv"  //libreria para convertir cadenas a numeros
	"strings"   //libreria para trabajar con cadenas en Go
)

func alerta(mensaje string) { //creamos la funcion alerta para mostrar mensaje en caso de error
	fmt.Println(mensaje)
}

func main() { //creamos la funcion main principal

	lectura := bufio.NewReader(os.Stdin) //creamos la variable que permite leer entradas por teclado

	fmt.Print("Ingrese el valor 1: ") //pedimos el primer valor al usuario
	primeraEntrada, _ := lectura.ReadString('\n') //se lee el texto ingresado en la consola
	float1, error := strconv.ParseFloat(strings.TrimSpace(primeraEntrada), 64) //aqui se convierte el valor ingresado como cadena, en numero float

	if error != nil {  //en esta condicion se maneja el error en caso de se presente una entrada no valida por el usuario
		fmt.Println(error)
		alerta("El primer valor debe ser un numero")
	}

	fmt.Println("Ingrese el valor 2: ")  //pedimos el segundo valor al usuario
	segundaEntrada, _ := lectura.ReadString('\n')  
	float2, error := strconv.ParseFloat(strings.TrimSpace(segundaEntrada), 64) //aqui se convierte el valor ingresado como cadena, en numero float

	if error != nil { //en esta condicion se maneja el error en caso de se presente una entrada no valida por el usuario
		fmt.Println(error)
		alerta("El segundo valor debe ser un numero")
	}

	suma := float1 + float2  //se suman los numeros
	suma = math.Round(suma*100) / 100 //redondeamos el numero multiplicandolo y dividiendolo por 100
	fmt.Printf("La suma de %v y %v es %v\n\n", float1, float2, suma)  //mostramos el resultado
}
