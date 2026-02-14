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

// BacklogPrioritario es el plan de crecimiento de GeoChat
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
func (c *CEO) EjecutarCicloDesarrollo() {
	c.Lock()
	defer c.Unlock()

	for i, f := range BacklogPrioritario {
		// La IA solo propone si el fondo del 15% (TokensGratis) tiene resto
		if f.Estado == "Pendiente" && c.TokensGratis >= f.Costo {

			// Restamos del presupuesto de crecimiento (Capa 4 Simulación)
			c.TokensGratis -= f.Costo

			// Generamos la propuesta usando los nombres exactos de types.go
			nueva := Propuesta{
				ID:            fmt.Sprintf("DEV-%d-%d", i, time.Now().Unix()),
				Modulo:        f.Nombre,
				Descripcion:   f.Descripcion,
				Monto:         f.Costo, // Antes era CostoTokens
				Impacto:       "Desarrollo de funcionalidad core",
				Estado:        "ESPERANDO_FIRMA", // Antes era Status
				Fecha:         time.Now(),
				RequiereFirma: true,
			}

			// Registramos la propuesta en el cerebro del CEO
			c.Propuestas = append(c.Propuestas, nueva)

			// Marcamos como 'En_Sandbox' para evitar duplicados
			BacklogPrioritario[i].Estado = "En_Sandbox"

			log.Printf("🤖 IA CEO: Propuesta generada para '%s'.", f.Nombre)

			// Notificación al Jefe (Usamos EnviarMensajeSoberano que definimos antes)
			texto := fmt.Sprintf("💡 *NUEVA PROPUESTA DE DESARROLLO*\n\n"+
				"Módulo: *%s*\n"+
				"Descripción: %s\n"+
				"Inversión: *%.1f PAXG*\n\n"+
				"Esperando su firma en el Panel de Control.",
				f.Nombre, f.Descripcion, f.Costo)

			err := c.EnviarMensajeSoberano(texto)
			if err != nil {
				log.Printf("⚠️ Error al notificar al Jefe: %v", err)
			}
		}
	}
}
