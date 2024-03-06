package files

import (
	"bufio"

	"os"

	"github.com/LuisDaga07/goDesde0/ejercicios"

	"fmt"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablaDeMultiplicar()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo!" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}
func SumaTabla() {
	var texto string = ejercicios.TablaDeMultiplicar()
	if !Append(filename, texto) { // El signo de interrogacion al principio !Append: quiere decir que es un condicional: Si no o es == false
		fmt.Println("Error al concatenar el archivo")
	}

}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al appendar el archivo" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
