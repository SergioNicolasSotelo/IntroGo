package main

import "fmt"

func main() {
	act1()
	act2()
}

func act1() {
	nombre := "Sergio Sotelo"
	direccion := "Calle falsa 123"

	fmt.Println(nombre, direccion)
}

func act2() {
	var temperatura float32
	var humedad int
	var presionAtmosferica float32

	temperatura = 23.0
	humedad = 49
	presionAtmosferica = 1009.4

	fmt.Println(temperatura, humedad, presionAtmosferica)
}
