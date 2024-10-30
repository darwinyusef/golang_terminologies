package main

import (
	"fmt"
)

// Estructura Banco
type Banco struct {
	saldo float64 // Campo privado
}

// Método para crear una nueva instancia de Banco
func NuevaCuenta() *Banco {
	return &Banco{saldo: 0} // Inicializa el saldo en 0
}

// Método público para depositar dinero
func (b *Banco) Depositar(cantidad float64) {
	if cantidad > 0 {
		b.saldo += cantidad
	}
}

// Método público para obtener el saldo
func (b *Banco) ObtenerSaldo() float64 {
	return b.saldo
}

// Método privado que intenta acceder al saldo
func (b *Banco) mostrarSaldoInterno() {
	fmt.Printf("Saldo interno: %.2f\n", b.saldo)
}

func main() {
	cuenta := NuevaCuenta()
	cuenta.Depositar(100)

	// Intentamos acceder a la variable privada directamente (esto no compilará)
	// fmt.Println(cuenta.saldo) // Esto generará un error de compilación

	// Usamos el método público para obtener el saldo
	fmt.Printf("Saldo actual: %.2f\n", cuenta.ObtenerSaldo())

	// Llamamos a un método que puede acceder al saldo privado
	cuenta.mostrarSaldoInterno()
}
