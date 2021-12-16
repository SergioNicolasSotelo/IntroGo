package main

import (
	"fmt"
)

func main() {
	longitudYLetras("palabra")
	aplicarDescuento(100, 50)
	darPrestamo(26, true, 2, 165000)
	queMesEs(11)
	var listaAlumnos = []string{"sergio", "tomas"}
	agregarALista(listaAlumnos, "bautista")
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	queEdadTiene(employees, "Benjamin")
	cuantosMayoresHay(employees)
	agregarPersona(employees, "Federico", 25)
	eliminarPersona(employees, "Pedro")
}

func longitudYLetras(palabra string) {
	fmt.Println("Longitud de palabra: ", len(palabra), " letras")

	for _, letra := range palabra {
		fmt.Println(string(letra))
	}
}

func aplicarDescuento(precio int32, descuento int32) {
	descontado := 100 - descuento
	fmt.Println(precio / 100 * descontado)
}

func darPrestamo(edad int32, empleado bool, antiguedad int32, salario int32) {
	if edad >= 22 && empleado == true && antiguedad >= 1 {
		if salario >= 100000 {
			fmt.Println("Puede obtener un préstamo sin interés.")
		} else {
			fmt.Println("Puede obtener un préstamo con interés.")
		}
	} else {
		fmt.Println("No puede obtener un préstamo.")
	}
}

func queMesEs(mes int32) {
	var nombreDeMes string
	switch mes {
	case 1:
		nombreDeMes = "Enero"
	case 2:
		nombreDeMes = "Febrero"
	case 3:
		nombreDeMes = "Marzo"
	case 4:
		nombreDeMes = "Abril"
	case 5:
		nombreDeMes = "Mayo"
	case 6:
		nombreDeMes = "Junio"
	case 7:
		nombreDeMes = "Julio"
	case 8:
		nombreDeMes = "Agosto"
	case 9:
		nombreDeMes = "Septiembre"
	case 10:
		nombreDeMes = "Octubre"
	case 11:
		nombreDeMes = "Noviembre"
	case 12:
		nombreDeMes = "Diciembre"
	default:
		nombreDeMes = "Ese no es un nombre de mes válido"
	}
	fmt.Println(mes, nombreDeMes)
}

func agregarALista(lista []string, alumno string) {
	lista = append(lista, alumno)
	fmt.Println(lista)
}

func queEdadTiene(personas map[string]int, nombre string) {
	fmt.Println(personas[nombre])
}

func cuantosMayoresHay(personas map[string]int) {
	var contador int
	for _, persona := range personas {
		if persona >= 21 {
			contador++
		}
	}
	fmt.Println(contador)
}

func agregarPersona(personas map[string]int, nombre string, edad int) {
	fmt.Println(personas)
	personas[nombre] = edad
	fmt.Println(personas)
}

func eliminarPersona(personas map[string]int, nombre string) {
	fmt.Println(personas)
	delete(personas, nombre)
	fmt.Println(personas)
}
