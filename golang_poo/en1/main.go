package main

import (
	"fmt"
)

// Estructura CuentaBancaria
type CuentaBancaria struct {
	saldo float64 // Campo privado
}

// Método para crear una nueva cuenta
func NuevaCuentaBancaria() *CuentaBancaria {
	return &CuentaBancaria{saldo: 0} // Inicializa el saldo en 0
}

// Método público para depositar dinero
func (c *CuentaBancaria) Depositar(cantidad float64) {
	if cantidad > 0 {
		c.saldo += cantidad
	}
}

// Método público para retirar dinero
func (c *CuentaBancaria) Retirar(cantidad float64) bool {
	if cantidad > 0 && cantidad <= c.saldo {
		c.saldo -= cantidad
		return true
	}
	return false
}

// Método público para obtener el saldo
func (c *CuentaBancaria) ObtenerSaldo() float64 {
	return c.saldo
}

func main() {
	cuenta := NuevaCuentaBancaria()
	cuenta.Depositar(100)
	fmt.Printf("Saldo después del depósito: %.2f\n", cuenta.ObtenerSaldo()) // Saldo: 100.00

	if cuenta.Retirar(50) {
		fmt.Println("Retiro exitoso.")
	} else {
		fmt.Println("Fondos insuficientes.")
	}

	fmt.Printf("Saldo después del retiro: %.2f\n", cuenta.ObtenerSaldo()) // Saldo: 50.00
}
