package ai_ceo

import (
	"fmt"
	"geochat/internal/database"
	"log"
	"time"
)

// GenerarPropuestaDiaria consolida la información estratégica para tu firma soberana.
func (c *CEO) GenerarPropuestaDiaria() Propuesta {
	c.RLock()
	defer c.RUnlock()

	// 1. Recolectamos la evolución técnica detectada (Ya no dará error)
	evoluciones := c.GetEvolucionesPendientes()
	numEvoluciones := len(evoluciones)

	// 2. Construimos el informe de impacto
	detalle := fmt.Sprintf("GeoChat procesó %d interacciones sociales. Hay %d evoluciones técnicas en espera.", 
		c.Stats.BuenaOndaCount, 
		numEvoluciones,
	)
	
	log.Println("📊 IA 5 analizando el día:", detalle)

	// 3. Creamos la estructura de la Propuesta (Campos alineados con types.go)
	propuesta := Propuesta{
		ID:            fmt.Sprintf("PROP-%d", time.Now().Unix()),
		Modulo:        "Evolución Diaria",
		Descripcion:   detalle,
		Monto:         c.FondoGas * 0.05, // Usamos el 5% del fondo operativo
		Impacto:       "Crecimiento del 15% basado en actividad social",
		Estado:        "ESPERANDO_FIRMA", // Coincide con tu regla de oro
		RequiereFirma: true,
		CreatedAt:     time.Now(),
	}

	// 4. PERSISTENCIA EN POSTGRES
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