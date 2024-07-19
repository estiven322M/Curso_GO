package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (mypc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco Duro y es una %s", mypc.ram, mypc.disk, mypc.brand)

}

func main() {
	mypc := pc{ram: 16, brand: "Lenovo", disk: 512}
	fmt.Println(mypc)
}
