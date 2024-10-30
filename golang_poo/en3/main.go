package main

import (
	"fmt"
)

// Estructura Banco
type Banco struct {
	saldo float64 // Campo privado
}

type Private struct {
	accept bool
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

// Método privado que no puede ser llamado desde fuera del paquete
func (b *Banco) calcularIntereses(private bool) float64 {
	p := &Private{}
	p.privates(private)
	return b.saldo * 0.05 // Ejemplo de cálculo de intereses
}

func (info *Private) privates(private bool) bool {
	info.accept = private
	if info.accept == true {
		return true
	}
	panic("¡El dato al cual usted desea ingresar es privado!")
}

// Método público que muestra el saldo total con intereses
func (b *Banco) MostrarSaldoConIntereses() {
	intereses := b.calcularIntereses(true) // Llama al método privado
	fmt.Printf("Saldo con intereses: %.2f\n", b.ObtenerSaldo()+intereses)
}

func main() {
	cuenta := NuevaCuenta()
	cuenta.Depositar(100)

	// Intentar acceder a la variable privada directamente generará un error de compilación
	fmt.Printf("Saldo actual: %.2f\n", cuenta.ObtenerSaldo()) // Saldo: 100.00

	// Mostrar el saldo con intereses usando un método público
	cuenta.MostrarSaldoConIntereses() // Saldo con intereses: 105.00

	// Intentar llamar a calcularIntereses directamente generará un error
	intereses := cuenta.calcularIntereses(true) // Esto generará un error de compilación
	defer fmt.Println(intereses)

	p := &Private{}

	// Prueba de la función
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de pánico:", r)
		}
	}()

	// Cambiamos el valor a true
	fmt.Println(p.privates(true)) // Salida: true

	// Esto provocará un pánico
	fmt.Println(p.privates(false)) // Esto provocará pánico
}
