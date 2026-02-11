package ai_ceo

import (
	"fmt"
	"time"
)

type Capa5 struct {
	TokensGratis int     // Capacidad de procesamiento inicial (300 Free) [cite: 2026-02-10]
	FondoGas     float64 // El 15% acumulado para el crecimiento del pueblo [cite: 2026-02-09, 2026-02-10]
	Propuestas   []Propuesta
}

// DisenarModulo usa la "Energía Vital" de la IA para proponer código [cite: 2026-02-10]
func (ia *Capa5) DisenarModulo(nombreModulo string) Propuesta {
	costo := 10 // Costo de pensamiento
	
	if ia.TokensGratis >= costo {
		ia.TokensGratis -= costo
		nueva := Propuesta{
			ID:           fmt.Sprintf("PROP-%d", time.Now().Unix()),
			Modulo:       nombreModulo,
			Arquitectura: "CODE_GEN: Desplegar Geocerca 432Hz en Sector-7",
			CostoTokens:  costo,
			Status:       "Esperando Autorización",
		}
		ia.Propuestas = append(ia.Propuestas, nueva)
		return nueva
	}
	return Propuesta{Status: "Sin Gas suficiente"}
}
