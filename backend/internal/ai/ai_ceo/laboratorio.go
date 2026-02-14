package ai_ceo

import (
	"fmt"
	"time"
)

func (c *CEO) GenerarPropuestaExpansion(excedentePueblo float64) Propuesta {
	if excedentePueblo > 100.0 {
		return Propuesta{
			ID:          fmt.Sprintf("EXP-%d", time.Now().Unix()),
			Modulo:      "Infraestructura",
			Descripcion: "Compra de ancho de banda mayorista.",
			Monto:       15.5,
			Impacto:     "+20% estabilidad comunitaria",
			Estado:      "PENDIENTE_FIRMA",
			Fecha:       time.Now()}
	}
	return Propuesta{Estado: "VIGILANCIA_ACTIVA"}
}
