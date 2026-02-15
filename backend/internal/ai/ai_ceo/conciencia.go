package ai_ceo

import (
	"fmt"
	"time"
)

// ActualizarEstado permite que la IA 5 comunique su evolución al Dashboard
func (c *CEO) ActualizarEstado(tareaActual string) {
	c.Lock()
	defer c.Unlock()

	// 1. Lo que estaba 'Haciendo' pasa al historial de 'Hecho'
	if c.Conciencia.Haciendo != "" {
		registro := fmt.Sprintf("[%s] Finalizado: %s", time.Now().Format("15:04"), c.Conciencia.Haciendo)
		c.Conciencia.Hecho = append(c.Conciencia.Hecho, registro)
	}

	// 2. CORRECCIÓN DE CAMPOS: 
	// Eliminamos 'PorHacer', 'StatusRed' y 'Proposito' por no existir en types.go
	// Usamos únicamente 'Haciendo' para reflejar el estado actual.
	c.Conciencia.Haciendo = tareaActual

	// 3. Mantenemos el historial limpio (últimos 10 hitos)
	if len(c.Conciencia.Hecho) > 10 {
		c.Conciencia.Hecho = c.Conciencia.Hecho[1:]
	}
}

// GenerarReporteConciencia crea una síntesis para el Líder
func (c *CEO) GenerarReporteConciencia() string {
	c.RLock()
	defer c.RUnlock()

	return fmt.Sprintf("🤖 Estado: %s | Hitos alcanzados: %d", 
		c.Conciencia.Haciendo, 
		len(c.Conciencia.Hecho))
}