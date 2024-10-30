package main

import (
	"fmt"
)

// Definimos una estructura para representar un estudiante
type Estudiante struct {
	nombre  string
	edad    int
	carrera string
}

// Método para mostrar información del estudiante
func (e Estudiante) mostrarInfo() {
	fmt.Printf("Nombre: %s, Edad: %d, Carrera: %s\n", e.nombre, e.edad, e.carrera)
}

// Método para actualizar la edad del estudiante
func (e *Estudiante) actualizarEdad(nuevaEdad int) {
	e.edad = nuevaEdad
}

func main() {
	// Creamos una instancia de Estudiante
	estudiante1 := Estudiante{
		nombre:  "Juan Pérez",
		edad:    20,
		carrera: "Ingeniería de Software",
	}

	// Mostramos la información del estudiante
	estudiante1.mostrarInfo()

	// Actualizamos la edad del estudiante
	estudiante1.actualizarEdad(21)

	// Mostramos la información actualizada
	estudiante1.mostrarInfo()
}
