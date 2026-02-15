package ai_ceo

import (
	"fmt"
	"log"
)

// EjecutarMovimiento recibe la propuesta y la firma del Líder para proceder.
// [cite: 2026-02-10] Sin tu firma validada, la IA no ejecuta ningún movimiento.
func (c *CEO) EjecutarMovimiento(p Propuesta, firmaRecibida string) {
	c.Lock()
	defer c.Unlock()

	// 1. VALIDACIÓN SOBERANA: Mi Firma es la Ley
	// Comparamos la firma enviada con la guardada en el Perfil del Líder
	if firmaRecibida == "" || firmaRecibida != c.Lider.FirmaDigital {
		log.Printf("🚨 INTENTO DE BRECHA SOBERANA: Se intentó ejecutar '%s' sin firma válida.", p.Modulo)
		return 
	}

	log.Printf("🛡️ Firma Soberana validada para el módulo: %s", p.Modulo)
	
	// 2. EVOLUCIÓN DEL SISTEMA
	// Llamamos a InyectarCodigoFuncional (que vive en ceo.go)
	c.InyectarCodigoFuncional(p) 

	// 3. DOCUMENTACIÓN EN VAULT [cite: 2026-02-11]
	// Usamos el método que YA ESTÁ declarado en capa5_vault.go
	c.DocumentarEnVault(fmt.Sprintf("Módulo %s ejecutado y verificado por el Líder. Firma: %s", p.Modulo, firmaRecibida))
	
	log.Printf("✅ [GEOCHAT-CEO]: Movimiento ejecutado y documentado con éxito.")
}

// NOTA: El método DocumentarEnVault ha sido eliminado de aquí porque 
// ya vive oficialmente en capa5_vault.go. Go lo encontrará automáticamente.