package ai_ceo

import (
	"time"
)

// EvaluarTemperatura: El filtro de "Cabeza Fría" de la IA 5 [cite: 2026-02-10]
func (c *CEO) EvaluarTemperatura(p Propuesta) (bool, string) {
	ahora := time.Now()
	
	// 1. Detectar contextos de "Cabeza Caliente" (Decisiones nocturnas)
	// Si el impacto financiero es alto y es de madrugada
	if p.ImpactoFinanciero > 10.0 && (ahora.Hour() > 22 || ahora.Hour() < 6) {
		return true, "🌙 Jefe, es tarde. Las decisiones de alto impacto (>10%) requieren luz de día y mente fresca. ¿Esperamos a mañana para firmar?"
	}

	// 2. Evaluar contra el "Perfil Filantrópico" [cite: 2026-02-02]
	if p.Status == "Modo_Extraccion" {
		alternativa := "🚀 Jefe, detecto un movimiento de extracción. ¿Qué tal si activamos el 'Modo Tesla' temporal? Esto potenciaría 200 nodos nuevos en Avellaneda."
		return true, alternativa
	}

	return false, ""
}