package main

import (
	"fmt"
	"time"

	"Estructuras/User"
)

type jhon struct {
	User.Usuario
}

func main() {
	u := new(jhon)
	u.AltaUsuario(1, "Jhon Mendez", time.Now(), true)
	fmt.Println(u.Usuario)

}
