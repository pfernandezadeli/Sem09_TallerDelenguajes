package tarifas

import "fmt"

// TipoServicio representa la modalidad de env?o.
type TipoServicio int

const (
	Estandar TipoServicio = iota
	Express
	Prioritario
)

const (
	TarifaBase = 8.00
	PrecioKg   = 2.50
	PrecioKm   = 0.15
	PesoMaximo = 50
	IGV        = 0.18
)

func CalcularSubtotal(pesoKg float64, distanciaKm int, servicio TipoServicio) (float64, error) {
	return 0, fmt.Errorf("sin implementar")
}

func AplicarIGV(subtotal float64) (igv, total float64) {
	return
}
