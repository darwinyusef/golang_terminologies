package main

import (
	"fmt"
	"math"
)

// Interfaz Figura
type Figura interface {
	Area() float64
}

// Estructura Circulo
type Circulo struct {
	radio float64
}

// Método para calcular el área del círculo
func (c Circulo) Area() float64 {
	return math.Pi * c.radio * c.radio
}

// Estructura Cuadrado
type Cuadrado struct {
	lado float64
}

// Método para calcular el área del cuadrado
func (q Cuadrado) Area() float64 {
	return q.lado * q.lado
}

// Función para imprimir el área de una figura
func imprimirArea(f Figura) {
	fmt.Printf("El área es: %.2f\n", f.Area())
}

func main() {
	circulo := Circulo{radio: 5}
	cuadrado := Cuadrado{lado: 4}

	imprimirArea(circulo)
	imprimirArea(cuadrado)
}
