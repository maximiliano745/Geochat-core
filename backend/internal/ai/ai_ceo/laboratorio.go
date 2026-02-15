package ai_ceo

import (
	"fmt"
	"time"
)

// GenerarPropuestaExpansion simula el crecimiento del 15% [cite: 2026-02-10]
func (c *CEO) GenerarPropuestaExpansion(excedentePueblo float64) Propuesta {
	
	// 1. Caso de Expansión Activa
	if excedentePueblo > 100.0 {
		return Propuesta{
			ID:            fmt.Sprintf("EXP-%d", time.Now().Unix()),
			Modulo:        "Infraestructura",
			Descripcion:   "Compra de ancho de banda mayorista para el pueblo.",
			Monto:         15.5,
			Impacto:       "+20% estabilidad comunitaria",
			Estado:        "PENDIENTE_FIRMA",
			RequiereFirma: true,
			CreatedAt:     time.Now(), // <--- CORREGIDO: Antes decía Fecha
		}
	} // <--- ESTA LLAVE FALTABA (Cierra el bloque IF)

	// 2. Caso de Vigilancia (Ahora sí es alcanzable)
	// Si no hay excedente suficiente, el CEO se queda en modo centinela.
	return Propuesta{
		ID:          "WATCH-001",
		Modulo:      "Vigilancia",
		Descripcion: "Escaneando excedentes insuficientes.",
		Estado:      "VIGILANCIA_ACTIVA",
		CreatedAt:   time.Now(),
	}
}