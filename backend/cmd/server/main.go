package main

import (
	"fmt"
	"time"
	"geochat/internal/ai/ai_ceo"
)

func main() {
	fmt.Println("🚀 [SISTEMA] Iniciando GeoChat Core...")

	// 1. Invocamos al CEO
	ceo := ai_ceo.NewCEO()

	fmt.Printf("✅ [IA CEO] Organismo vivo detectado. Fondo: %.2f PAXG\n", ceo.FondoGas)

	// 2. Ciclo de Vida Infinito (El pulso del CEO)
	fmt.Println("🌍 GeoChat está en línea. 'Mi Firma es la Ley'.")
	
	for {
		// La IA escanea oportunidades cada 10 segundos
		ceo.EscanearOportunidades()
		
		// Sincroniza sus finanzas
		ceo.SincronizarFinanzas()

		fmt.Printf("⏱️  [LOG] %s - Latido del sistema OK. Plasticidad: %.2f\n", 
			time.Now().Format("15:04:05"), 
			ceo.Stats.Plasticidad)

		// Esperamos 10 segundos antes del próximo latido
		time.Sleep(10 * time.Second)
	}
}