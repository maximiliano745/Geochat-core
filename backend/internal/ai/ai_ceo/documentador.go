package ai_ceo

import (
	"fmt"
	"log"
	"time"
)

// RegistrarEvolucion es el método que guarda cada paso del Proyecto GeoChat.
// Ahora usa directamente el puntero del CEO para notificar.
func (c *CEO) RegistrarEvolucion(modulo string, detalle string) {
	c.Lock()
	defer c.Unlock()

	// 1. Crear la entrada para el Vault Digital [cite: 2026-02-11]
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entrada := fmt.Sprintf("[%s] EVOLUCIÓN: %s | DETALLE: %s", timestamp, modulo, detalle)

	// 2. Guardar en el historial del Líder (el Vault del CEO)
	c.Lider.HistorialFirma = append(c.Lider.HistorialFirma, entrada)
	
	log.Printf("📂 Vault: Documento registrado - %s", modulo)

	// 3. ✅ CORRECCIÓN: Usar el método del CEO para notificar al Jefe
	// Antes aquí buscabas a 'WhatsAppService', ahora el CEO habla solo.
	err := c.EnviarMensajeSoberano(fmt.Sprintf("📂 *GEOCHAT VAULT*\nNueva evolución registrada: %s\nDetalle: %s", modulo, detalle))
	
	if err != nil {
		log.Printf("⚠️ Error al notificar al Jefe: %v", err)
	}
}