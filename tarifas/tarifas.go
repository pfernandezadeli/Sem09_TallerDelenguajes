package tarifas

import "fmt"

// TipoServicio representa la modalidad de envío.
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

// CalcularSubtotal devuelve el costo sin IGV o un error si los datos no son válidos.
func CalcularSubtotal(pesoKg float64, distanciaKm int, servicio TipoServicio) (float64, error) {
	if pesoKg <= 0 || pesoKg > PesoMaximo {
		return 0, fmt.Errorf("peso no válido: %.2f kg (debe estar entre 0 y %d kg)", pesoKg, PesoMaximo)
	}
	if distanciaKm <= 0 {
		return 0, fmt.Errorf("distancia no válida: %d km (debe ser mayor que 0 km)", distanciaKm)
	}
	var recargo float64
	switch servicio {
	case Estandar:
		recargo = 1.0
	case Express:
		recargo = 1.3
	case Prioritario:
		recargo = 1.6
	default:
		return 0, fmt.Errorf("servicio no válido: %d", servicio)
	}
	subtotal := (TarifaBase + pesoKg*PrecioKg + float64(distanciaKm)*PrecioKm) * recargo
	return subtotal, nil
}

// AplicarIGV devuelve el IGV y el total usando retornos con nombre.
func AplicarIGV(subtotal float64) (igv, total float64) {
	igv = subtotal * IGV
	total = subtotal + igv
	return
}
