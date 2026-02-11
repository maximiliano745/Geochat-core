package business

import (
	"fmt"
)

// EconomicHealth define los parámetros de Salud Económica de una zona específica.
type EconomicHealth struct {
	InflationRate     float64 // Tasa de inflación local (ej: 20.0 para 20%)
	BandwidthScarcity float64 // Escasez de internet (0.0 a 1.0)
	AdRevenueYield    float64 // Rendimiento de IA 1 en la zona
}

// CalculateDynamicPrice es el Algoritmo de Equidad "Pueblo para el Pueblo".
// Ajusta el costo de los servicios para que nadie quede fuera del sistema por motivos económicos.
func CalculateDynamicPrice(basePrice float64, health EconomicHealth) float64 {
	// 1. Ajuste por Inflación: Si sube la inflación, baja el precio relativo para proteger al usuario.
	adjustment := 1.0 - (health.InflationRate / 100)
	
	// 2. Incentivo de Infraestructura: Si no hay internet, el Nodo es un 50% más barato para incentivar la red.
	if health.BandwidthScarcity > 0.7 {
		adjustment *= 0.5 
	}

	// 3. Subsidio por Publicidad: Si la IA 1 genera ingresos, bajamos el costo al usuario.
	if health.AdRevenueYield > 5.0 {
		adjustment -= 0.1
	}

	// Límite de seguridad Soberana: Nunca menos del 30% del costo operativo base para no descapitalizar.
	finalPrice := basePrice * adjustment
	if finalPrice < (basePrice * 0.3) {
		finalPrice = basePrice * 0.3
	}

	return finalPrice
}

// CEOMovement representa un movimiento financiero de la Capa 5.
type CEOMovement struct {
	TotalAmount     float64
	GrowthFund15    float64 // El 15% destinado al crecimiento del pueblo [cite: 2026-02-10]
	Infrastructure  float64 // Inversión en hardware/ancho de banda
	ValidationKey   string  // "Mi Firma es la Ley" [cite: 2026-02-10]
}

// PrepareTransaction separa el 15% para el fondo de crecimiento de forma automática.
func (ceo *CEOMovement) PrepareTransaction(payment float64) {
	ceo.TotalAmount = payment
	ceo.GrowthFund15 = payment * 0.15
	fmt.Printf("IA 5: Calculando equidad... Reservando %.2f PAXG para el fondo de crecimiento.\n", ceo.GrowthFund15)
}