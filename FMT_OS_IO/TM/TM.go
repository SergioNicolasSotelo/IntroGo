package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	guardarEnArchivo()
	leerArchivo()
}

func guardarEnArchivo() {
	var id int
	var precio float64
	var cantidad int

	fmt.Println("Introducir id:")
	fmt.Scanln(id)
	fmt.Scanln(&id)
	fmt.Println("Introducir precio:")
	fmt.Scanln(precio)
	fmt.Scanln(&precio)
	fmt.Println("Introducir cantidad:")
	fmt.Scanln(cantidad)
	fmt.Scanln(&cantidad)

	res := fmt.Sprintf("\n %d \t %f \t %d", id, precio, cantidad)
	d1 := []byte(res)
	data := append(d1)
	err := os.WriteFile("./DatosProductos.txt", data, 0644)
	fmt.Println(err)
}

func leerArchivo() {
	data, err := os.ReadFile("./DatosProductos.txt")

	if data != nil {
		aux := string(data)
		lineaPorLinea := strings.Split(aux, "\n")
		for _, linea := range lineaPorLinea {
			fmt.Println(linea)
		}
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Archivo sin entradas.")
	}
}
