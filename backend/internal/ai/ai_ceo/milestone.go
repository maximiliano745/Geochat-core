package ai_ceo

import (
	"log"
	"os"
	"path/filepath"
)

// InyectarHitoCero genera la primera propuesta real de código
func InyectarHitoCero() Propuesta { // <-- Cambiado Proposal por Propuesta
	// 1. Definimos la propuesta estratégica alineada con el ADN
	p := Propuesta{
		ID:                "HITO-000-START",
		Modulo:            "Gateway PAXG v1", // Usamos Modulo en lugar de Description
		Monto:             0.0,
		Tipo:              "LOGIC",
		Status:            "PENDIENTE_FIRMA", // Estado consistente
		RequiereFirma:     true,
	}

	// 2. Simulamos la redacción del código por la IA
	codigoFuente := `
package core
import "log"
func InitGatewayPAXG() {
    log.Println("✨ [CORE] Gateway PAXG v1 Operativo: El Oro fluye por GeoChat.")
}`

	// 3. Depositamos el código en el Laboratorio para tu revisión
	wd, _ := os.Getwd()
	labPath := filepath.Join(wd, "lab", "gateway_paxg_v1.go")
	
	// Crear carpeta lab si no existe
	os.MkdirAll(filepath.Dir(labPath), 0755)

	err := os.WriteFile(labPath, []byte(codigoFuente), 0644)
	if err != nil {
		log.Printf("❌ Error al depositar en Lab: %v", err)
	}

	log.Printf("🧪 Hito Cero: Código depositado en %s. Esperando tu firma...", labPath)
	return p
}

// FirmarYDesplegar: Aquí "Tu Firma es la Ley" [cite: 2026-02-10]
func FirmarYDesplegar(p Propuesta, firma string) bool { // <-- Cambiado Proposal por Propuesta
	if firma != "MI_FIRMA_ES_LA_LEY" {
		log.Println("🚫 Firma inválida. El código permanece en el laboratorio.")
		return false
	}

	// MOVEMOS EL CÓDIGO DEL LAB AL CORE
	wd, _ := os.Getwd()
	origen := filepath.Join(wd, "lab", "gateway_paxg_v1.go")
	destino := filepath.Join(wd, "internal", "core", "gateway_paxg_v1.go")

	os.MkdirAll(filepath.Dir(destino), 0755)

	err := os.Rename(origen, destino)
	if err != nil {
		log.Printf("❌ Error en el despliegue: %v", err)
		return false
	}

	log.Println("🔥 [SISTEMA] ¡Hito Cero Desplegado! El Watcher debería reiniciar GeoChat ahora.")
	return true
}