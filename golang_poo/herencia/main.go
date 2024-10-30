package main

import (
	"fmt"
)

// Definimos una estructura base llamada "Animal"
type Animal struct {
	Name string
}

// Método para el tipo Animal
func (a Animal) Speak(raza string) string {
	return "Animal tipo: " + raza + " hace un sonido"
}

// Definimos una estructura que "hereda" de Animal
type Type struct {
	Animal // Composición: Type tiene un Animal
	Breed  string
}

// Método específico para Type
func (d Type) Speak() string {
	return "Mi nombre es " + d.Name
}

func main() {
	// Creamos un perro
	myDog := Type{
		Animal: Animal{Name: "Fido"},
		Breed:  "Labrador",
	}

	// Llamamos a los métodos
	fmt.Println(myDog.Speak())
	fmt.Println(myDog.Animal.Speak("Perro"))

	// Creamos un gato
	myCat := Type{
		Animal: Animal{Name: "Pepito"},
		Breed:  "Angora",
	}

	fmt.Println(myCat.Speak())
	fmt.Println(myCat.Animal.Speak("Gato"))
}
