package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngresarNumero() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err := strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				for i := 1; i <= 10; i++ {
					fmt.Printf("%d x %d = %d \n", numero, i, i*numero)
				}
				break
			}
		}
	}

}
