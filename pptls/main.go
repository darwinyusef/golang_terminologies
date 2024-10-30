package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Constantes para las opciones
const (
	Piedra  = "piedra"
	Papel   = "papel"
	Tijera  = "tijera"
	Lagarto = "lagarto"
	Spock   = "spock"
)

// Función para determinar el ganador
func determinarGanador(opcionJugador string, opcionComputador string) string {
	if opcionJugador == opcionComputador {
		return "Es un empate"
	}

	switch opcionJugador {
	case Piedra:
		if opcionComputador == Tijera || opcionComputador == Lagarto {
			return "Jugador gana"
		}
	case Papel:
		if opcionComputador == Piedra || opcionComputador == Spock {
			return "Jugador gana"
		}
	case Tijera:
		if opcionComputador == Papel || opcionComputador == Lagarto {
			return "Jugador gana"
		}
	case Lagarto:
		if opcionComputador == Spock || opcionComputador == Papel {
			return "Jugador gana"
		}
	case Spock:
		if opcionComputador == Tijera || opcionComputador == Piedra {
			return "Jugador gana"
		}
	}
	return "Computador gana"
}

// Función para obtener la elección aleatoria de la computadora
func opcionAleatoria() string {
	opciones := []string{Piedra, Papel, Tijera, Lagarto, Spock}
	rand.Seed(time.Now().UnixNano())
	return opciones[rand.Intn(len(opciones))]
}

func main() {
	fmt.Println("Bienvenido al juego Piedra, Papel, Tijera, Lagarto, Spock")

	// Jugador elige
	var opcionJugador string
	fmt.Println("Elige una opción en minuscula:")
	fmt.Scanln(&opcionJugador)

	// Convertir la elección del jugador a capitalizado
	switch opcionJugador {
	case "piedra":
		opcionJugador = Piedra
	case "papel":
		opcionJugador = Papel
	case "tijera":
		opcionJugador = Tijera
	case "lagarto":
		opcionJugador = Lagarto
	case "spock":
		opcionJugador = Spock
	default:
		fmt.Println("Opción no válida. Inténtalo de nuevo.")
		return
	}

	// Computadora elige
	opcionComputador := opcionAleatoria()
	fmt.Printf("La computadora eligió: %s\n", opcionComputador)

	// Determinar y mostrar el ganador
	resultado := determinarGanador(opcionJugador, opcionComputador)
	fmt.Println(resultado)
}
