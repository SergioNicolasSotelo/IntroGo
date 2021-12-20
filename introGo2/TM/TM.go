package main

import (
	"errors"
	"fmt"
)

func main() {

}

func calcularImpuestos(salario int) int {
	if salario >= 150000 {
		return salario/100*17 + salario/100*10
	} else if salario >= 50000 && salario <= 150000 {
		return salario / 100 * 17
	} else {
		return 0
	}
}

func calcularPromedios(calificaciones ...int) int {
	var suma int

	for _, c := range calificaciones {
		if c < 0 {
			fmt.Println("Error: Hay un numero negativo entre las calificaciones.")
			return 0
		}
		suma += c
	}

	return suma / len(calificaciones)
}

func calcularSalario(minutos int, categoria string) int {
	var salario int
	if categoria == "c" {
		salario = (minutos / 60) * 1000
		return salario
	} else if categoria == "b" {
		salario = (minutos / 60) * 1500
		return salario + (salario / 100 * 20)
	} else {
		salario = (minutos / 60) * 3000
		return salario + (salario / 100 * 50)
	}
}

func calcularEstadisticas(operacion string) func(calificaciones ...int) int {
	switch operacion {
	case "minimo":
		return func(calificaciones ...int) int {
			min := calificaciones[0]
			for _, c := range calificaciones {
				if c < min {
					min = c
				}
			}
			return min
		}
	case "maximo":
		return func(calificaciones ...int) int {
			max := 0
			for _, c := range calificaciones {
				if c > max {
					max = c
				}
			}
			return max
		}
	case "promedio":
		return func(calificaciones ...int) int {
			var suma int
			for _, c := range calificaciones {
				suma += c
			}
			return suma / len(calificaciones)
		}
	default:
		return func(calificaciones ...int) int {
			fmt.Println("Error: No existe la operación solicitada.")
			return 0
		}
	}
}

func calcularAlimento(animal string) (func(cantidad int) float64, error) {
	switch animal {
	case "perro":
		return alimentarPerro, nil
	case "gato":
		return alimentarGato, nil
	case "hamster":
		return alimentarHamster, nil
	case "tarantula":
		return alimentarTarantula, nil
	default:
		return nil, errors.New("Error: No existe ese animal.")
	}
}

func alimentarPerro(cantidad int) float64 {
	return float64(cantidad * 10)
}
func alimentarGato(cantidad int) float64 {
	return float64(cantidad * 5)
}
func alimentarHamster(cantidad int) float64 {
	return float64(cantidad) * 0.25
}
func alimentarTarantula(cantidad int) float64 {
	return float64(cantidad) * 0.15
}
