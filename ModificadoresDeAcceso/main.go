package main

import (
	ModificadoresDeAcceso "ModificadoresDeAcceso/ModificadoresAcceso"
)

type Stivenson struct {
}

func main() {
	var myCar ModificadoresDeAcceso.CarPublic
	myCar.brand = "Toyota"
	myCar.year = 2023

}
