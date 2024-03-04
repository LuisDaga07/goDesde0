package main

import (
	"fmt"

	"github.com/LuisDaga07/goDesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1544)
	fmt.Println(estado)
	fmt.Println(texto)
}
