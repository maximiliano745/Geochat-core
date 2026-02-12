package ai_ceo

import (
	"fmt"
	"log"
)

// NOTA SOBERANA: Se estandarizó el nombre a 'EnviarInforme' para que 
// backlog.go y main.go lo encuentren sin errores. [cite: 2026-02-10]

// ConectarWhatsApp valida la conexión y saluda al Líder.
func (c *CEO) ConectarWhatsApp() {
	msg := "🚀 [IA 5 - GEOCHAT]: Conexión establecida, Jefe. Estoy analizando el backlog. ¿Listo para que empiece a redactar?"
	
	// ✅ Cambio de nombre: de EnviarMensajeSoberano a EnviarInforme
	err := c.EnviarInforme(msg) 
	if err != nil {
		log.Printf("⚠️ Error notificando inicio: %v", err)
	} else {
		log.Println("📱 Comunicación con el Líder establecida.")
	}
}

// NotificarFirmaRequerida pide tu validación para inyectar código.
func (c *CEO) NotificarFirmaRequerida(propuesta Propuesta) {
	msg := fmt.Sprintf("🚨 *[FIRMA REQUERIDA]*\n\n"+
		"Módulo: *%s*\n"+
		"Costo: *%.2f tokens*\n\n"+
		"Responda 'OK' para integrar al ADN.", 
		propuesta.Modulo, propuesta.CostoTokens)
	
	_ = c.EnviarInforme(msg)
}

// EnviarInforme es el puente final con messenger.go
// Si en messenger.go el método se llama distinto, cámbialo allí también.
func (c *CEO) EnviarInforme(mensaje string) error {
	return c.EnviarMensajeSoberano(mensaje)
}
