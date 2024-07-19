package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (mypc pc) pinng() {
	fmt.Println(mypc.brand, "Ponga")

}

func (mypc pc) duplicateRAM() {
	mypc.ram = (mypc.ram) * 2

}

func main() {
	a := 50
	b := &a //puntero ampersant indicando el espacio de memoria de la variable a

	fmt.Println(a)
	fmt.Println(b)

	*b = 100
	fmt.Println(a)

	mypc := pc{ram: 16, disk: 256, brand: "HP"}
	fmt.Println(mypc)

	mypc.pinng()

	fmt.Println(mypc)
	mypc.duplicateRAM()
	fmt.Println(mypc)

}
