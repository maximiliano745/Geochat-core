package ai_ceo

import (
	"fmt"
	"log"
	"time"
)

// Funcionalidad representa una meta en el Roadmap soberano
type Funcionalidad struct {
	Nombre      string
	Descripcion string
	Estado      string
	Costo       float64
}

// BacklogPrioritario es el plan de crecimiento de GeoChat [cite: 2026-02-09]
var BacklogPrioritario = []Funcionalidad{
	{
		Nombre:      "Native Red & Mapa",
		Descripcion: "Sincronización GPS real-time y malla Mesh soberana",
		Estado:      "Pendiente",
		Costo:       2.0,
	},
	{
		Nombre:      "Geocercas & Biometría",
		Descripcion: "POIs dinámicos y validación vocal 2FA para el Vault",
		Estado:      "Pendiente",
		Costo:       3.0,
	},
	{
		Nombre:      "Mercado P2P",
		Descripcion: "Intercambio de energía e internet sin intermediarios",
		Estado:      "Pendiente",
		Costo:       4.0,
	},
}

// EjecutarCicloDesarrollo es el motor autónomo de la IA 5.
// Transforma items del backlog en propuestas reales para el Líder.
func (c *CEO) EjecutarCicloDesarrollo() {
	c.Lock()
	defer c.Unlock()

	for i, f := range BacklogPrioritario {
		// La IA solo propone si el 15% tiene fondos (TokensGratis) [cite: 2026-02-10]
		if f.Estado == "Pendiente" && c.TokensGratis >= f.Costo {
			
			// Restamos del presupuesto de crecimiento de la IA
			c.TokensGratis -= f.Costo

			// Generamos la propuesta técnica soberana
			nueva := Propuesta{
				ID:                fmt.Sprintf("DEV-%d-%d", i, time.Now().Unix()),
				Modulo:            f.Nombre,
				Tipo:              "DESARROLLO",
				CostoTokens:       f.Costo,
				Status:            "ESPERANDO_FIRMA",
				RequiereFirma:     true,
				ImpactoFinanciero: 0.15,
			}

			// Registramos la propuesta en el cerebro del CEO
			c.Propuestas = append(c.Propuestas, nueva)
			
			// Marcamos como 'En_Sandbox' para que no se repita en el próximo tick
			BacklogPrioritario[i].Estado = "En_Sandbox"

			log.Printf("🤖 IA CEO: Propuesta generada para '%s'.", f.Nombre)

			// Notificación inmediata al Jefe
			texto := fmt.Sprintf("💡 *NUEVA PROPUESTA DE DESARROLLO*\n\n"+
				"Módulo: *%s*\n"+
				"Descripción: %s\n"+
				"Inversión: *%.1f tokens*\n\n"+
				"Responda 'OK' para iniciar la inyección de código.",
				f.Nombre, f.Descripcion, f.Costo)

			// Usamos el método oficial EnviarInforme
			err := c.EnviarInforme(texto)
			if err != nil {
				log.Printf("⚠️ Error al notificar al Jefe: %v", err)
			}
		}
	}
}