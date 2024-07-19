package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Porshe", year: 2020}
	fmt.Println(myCar)

	//otra manera de instanciar un structs
	var otherCar car
	otherCar.brand = "Citroen"
	otherCar.year = 1995

}
