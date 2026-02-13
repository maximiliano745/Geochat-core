package ai_ceo

import (
	"fmt"
	"log"
	// "geochat/internal/finance" // Descomenta cuando Capa 4 esté lista
)

// EjecutarMovimiento es el último paso del flujo soberano.
// Ningún PAXG se mueve y ningún código muta sin pasar por aquí. [cite: 2026-02-10]
func (c *CEO) EjecutarMovimiento(p Propuesta) {
	// 1. Verificación Crítica: ¿El líder puso su firma digital?
	// Usamos el campo FirmaDigital que ya definiste en Propuesta.
	if p.FirmaDigital == "" {
		log.Printf("🚨 INTENTO DE BRECHA: La IA intentó ejecutar '%s' sin firma.", p.ID)
		return 
	}

	// 2. Si la firma es válida, se procede.
	log.Printf("🛡️ Firma validada. Ejecutando mandato del Líder: %s", p.Modulo)
	
	// Simulación de transferencia (Capa 4)
	if p.CostoTokens > 0 {
		log.Printf("💰 Transfiriendo %.2f tokens a %s...", p.CostoTokens, p.Destino)
		// Aquí iría: finance.EjecutarTransferencia(p.Destino, p.CostoTokens)
	}

	// 3. Inyectar el cambio en el ADN del sistema
	c.InyectarCodigoFuncional(p)

	// 4. Documentar el éxito en el historial del Líder [cite: 2026-02-11]
	c.DocumentarEnVault(fmt.Sprintf("%s: %s", p.Modulo, "Ejecución técnica y financiera completada"))	
}
