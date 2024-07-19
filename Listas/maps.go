package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pedro"] = 25
	m["Paola"] = 23
	m["Ericka"] = 18

	fmt.Println(m)

	//recorrer el mapa
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar valor
	value := m["Jose"]
	fmt.Println(value)
}
