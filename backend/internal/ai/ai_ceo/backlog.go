package ai_ceo

import (
	"fmt"
	"log"
)

type Funcionalidad struct {
	Nombre      string
	Descripcion string
	Estado      string // "Pendiente", "En_Sandbox", "Listo_Para_Firma"
	Costo       int    
}

// Mapa de Ruta: 100 Core, 100 Confianza, 100 Economía [cite: 2026-02-10]
var BacklogPrioritario = []Funcionalidad{
	{"Native Red & Mapa", "Sincronización GPS real-time", "Pendiente", 100},
	{"Geocercas & Biometría", "POIs y 2FA Vocal", "Pendiente", 100},
	{"Mercado P2P", "Billetera e Intercambio", "Pendiente", 100},
}

func (c *CEO) EjecutarCicloDesarrollo() {
	for i, f := range BacklogPrioritario {
		if f.Estado == "Pendiente" && c.TokensGratis >= f.Costo {
			c.TokensGratis -= f.Costo
			nueva := Propuesta{
				ID:           fmt.Sprintf("DEV-%d", i),
				Modulo:       f.Nombre,
				Arquitectura: fmt.Sprintf("// Código para: %s", f.Descripcion),
				CostoTokens:  f.Costo,
				Status:       "Esperando Autorización",
			}
			c.Propuestas = append(c.Propuestas, nueva)
			BacklogPrioritario[i].Estado = "En_Sandbox"
			log.Printf("CEO: Generando propuesta para %s [Costo: %d]\n", f.Nombre, f.Costo)
		}
	}
}