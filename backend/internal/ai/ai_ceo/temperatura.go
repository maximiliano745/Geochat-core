package ai_ceo

import (
	"time"
)

// EvaluarTemperatura: El filtro de "Cabeza Fría" de la IA 5 [cite: 2026-02-10]
// Evita que el Líder tome decisiones impulsivas en momentos de fatiga.
func (c *CEO) EvaluarTemperatura(p Propuesta) (bool, string) {
	ahora := time.Now()
	
	// 1. Detectar contextos de "Cabeza Caliente" (Decisiones nocturnas)
	// Comparamos p.Monto contra el límite de impacto (>10 PAXG).
	// [Alineado con types.go: usamos p.Monto]
	if p.Monto > 10.0 && (ahora.Hour() > 22 || ahora.Hour() < 6) {
		return true, "🌙 Jefe, es tarde. Las decisiones de alto impacto (>10 PAXG) requieren luz de día y mente fresca. ¿Esperamos a mañana para firmar?"
	}

	// 2. Evaluar contra el "Perfil Filantrópico" [cite: 2026-02-02]
	// GeoChat es del pueblo: sugerimos el 'Modo Tesla' si el estado es de extracción.
	// [Alineado con types.go: usamos p.Estado]
	if p.Estado == "Modo_Extraccion" {
		alternativa := "🚀 Jefe, detecto un movimiento de extracción. ¿Qué tal si activamos el 'Modo Tesla' manualmente? Esto potenciaría 200 nodos nuevos en Avellaneda."
		return true, alternativa
	}

	return false, ""
}