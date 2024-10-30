package main

import (
	"fmt"
	"sync"
)

// Definimos un struct llamado 'Persona'
type Persona struct {
	Nombre string
	Edad   int
}

// Método para presentar a la persona
func (p Persona) Presentar() string {
	return fmt.Sprintf("Hola, soy %s y tengo %d años.", p.Nombre, p.Edad)
}

// Método para celebrar cumpleaños
func (p *Persona) CumplirAnios() {
	p.Edad++
}

// GestorDePersonas actúa como Singleton
type GestorDePersonas struct {
	persona Persona
}

var (
	instance *GestorDePersonas
	once     sync.Once
)

// ObtenerInstancia devuelve la instancia única del GestorDePersonas
func ObtenerInstancia(nombre string, edad int) *GestorDePersonas {
	once.Do(func() {
		instance = &GestorDePersonas{
			persona: Persona{Nombre: nombre, Edad: edad},
		}
	})
	return instance
}

// ActualizarPersona permite cambiar la persona gestionada
func (g *GestorDePersonas) ActualizarPersona(nombre string, edad int) {
	g.persona.Nombre = nombre
	g.persona.Edad = edad
}

func main() {
	// Obtenemos la única instancia del GestorDePersonas
	gestor := ObtenerInstancia("Juan", 30)

	// Presentamos a Juan
	fmt.Println(gestor.persona.Presentar())

	// Juan cumple años
	gestor.persona.CumplirAnios()

	// Presentamos nuevamente a Juan
	fmt.Println(gestor.persona.Presentar())

	gestorOtro := ObtenerInstancia("Juan", 30)
	fmt.Println(gestorOtro.persona.Presentar())

	// Intentamos actualizar la persona a Maria
	gestor.ActualizarPersona("Maria", 25)

	// Presentamos nuevamente, ahora debe ser Maria
	fmt.Println(gestor.persona.Presentar())
}
