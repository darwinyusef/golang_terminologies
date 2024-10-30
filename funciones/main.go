package main

import (
	"fmt"
)

func main() {
	fmt.Println("¡Hola mundo!")
	// Llama a otra función de este paquete.
	másAllaDelHola()
}

// Las funciones llevan parámetros entre paréntesis.
// Si no hay parámetros, los paréntesis siguen siendo obligatorios.
func másAllaDelHola() {
	fmt.Print("mas alla de la Luz")
}
