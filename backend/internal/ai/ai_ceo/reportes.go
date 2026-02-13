package ai_ceo

import (
	"fmt"
	"geochat/internal/database"
	"log"
	"time"
)

// GenerarPropuestaDiaria consolida la información estratégica para tu firma soberana.
// Esta función es el corazón del crecimiento del 15%. [cite: 2026-02-10]
func (c *CEO) GenerarPropuestaDiaria() Propuesta {
	c.RLock()
	defer c.RUnlock()

	// 1. Recolectamos la evolución técnica detectada por la IA
	evoluciones := c.GetEvolucionesPendientes()
	numEvoluciones := len(evoluciones)

	// 2. Construimos el informe de impacto
	detalle := fmt.Sprintf("GeoChat procesó %d interacciones sociales. Hay %d evoluciones técnicas en espera.", 
		c.Stats.BuenaOndaCount, 
		numEvoluciones,
	)
	
	log.Println("📊 IA 5 analizando el día:", detalle)

	// 3. Creamos la estructura de la Propuesta
	propuesta := Propuesta{
		ID:                fmt.Sprintf("PROP-%d", time.Now().Unix()),
		Modulo:            "Evolución Diaria",
		Monto:             c.FondoGas * 0.05, // Usamos el 5% del fondo operativo
		ImpactoFinanciero: 5.0,
		Status:            "PENDIENTE_FIRMA",
		RequiereFirma:     true,
		Tipo:              "Social/Técnico",
	}

	// 4. PERSISTENCIA EN POSTGRES: Para que aparezca en el AdminView automáticamente
	if database.DB != nil {
		query := `INSERT INTO tasks (titulo, descripcion, estado) 
				  VALUES ($1, $2, 'propuesta_ceo')`
		
		tituloPropuesta := fmt.Sprintf("Ciclo de Evolución: %s", propuesta.Modulo)
		_, err := database.DB.Exec(query, tituloPropuesta, detalle)
		
		if err != nil {
			log.Printf("⚠️ IA 5 Error al persistir propuesta diaria: %v", err)
		} else {
			log.Println("✅ IA 5: Propuesta diaria enviada a la cola de autorización del Soberano.")
		}
	}

	return propuesta
}