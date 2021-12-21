package main

import (
	"errors"
	"fmt"
)

func main() {

}

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

func impuestos1(salary int) (string, error) {
	if salary < 150000 {
		return "", &myError{message: "error: el salario ingresado no alcanza el mínimo imponible"}
	} else {
		return "Debe pagar impuestos.", nil
	}
}

/*Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/
func impuestos2(salary int) (string, error) {
	if salary < 150000 {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	} else {
		return "Debe pagar impuesto", nil
	}
}

/*
Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/
func impuestos3(salary int) (string, error) {
	if salary < 150000 {
		return "", fmt.Errorf("El salario ingresado, %d, no alcanza el mínimo imponible, que es de 150000.", salary)
	} else {
		return "Debe pagar impuesto", nil
	}
}

/*
 Ejercicio 3 - Impuestos de salario #3
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/
func calcularSalario(horasMes int, valorHora float64) (string, error) {
	salary := float64(horasMes) * valorHora
	if salary < 150000 {
		return "", fmt.Errorf("El salario ingresado, %f, no alcanza el mínimo imponible, que es de 150000.", salary)
	} else {
		return "Debe pagar impuesto", nil
	}
}

/*
Ejercicio 4 -  Impuestos de salario #4
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
Desarrolla las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
*/
