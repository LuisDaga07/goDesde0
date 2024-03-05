package ejercicios

import (
	
	"strconv"
)

func ConvertToInt(valor string) (int, string) {
	numero, err := strconv.Atoi(valor)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
		
	}
	
}
